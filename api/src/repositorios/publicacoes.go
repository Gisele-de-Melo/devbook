package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação n banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// BuscarPorID traz uma unica publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select p.*, u.Nick
		   from publicacoes as p inner join 
		        usuarios as u on p.autor_id = u.id
		   where p.id = ?`,
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhass, erro := repositorio.db.Query(
		`select distinct p.*, u.Nick
		   from publicacoes as p inner join 
		        usuarios as u on p.autor_id = u.id inner join
				seguidores as s on p.autor_id = s.usuario_id
		   where u.id = ? or s.seguidor_id = ?
		   order by p.criadaEm desc
		`,
		usuarioID,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhass.Close()

	var publicacoes []modelos.Publicacao

	for linhass.Next() {
		var publicacao modelos.Publicacao
		if erro = linhass.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	if erro = linhass.Err(); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes 
		  set titulo = ?
    	    , conteudo = ?
		 where id = ?	
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação no banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`delete from publicacoes 
		 where id = ?	
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario traz as publicações de um usuário em específico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select p.*, u.Nick
		   from publicacoes as p inner join 
		        usuarios as u on p.autor_id = u.id
		   where p.autor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes 
		    set curtidas = curtidas + 1
		  where id = ?	
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}

// Descurtir remove uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(
		`update publicacoes 
		    set curtidas = curtidas - 1
		  where id = ?	
		    and curtidas > 0
		`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoID)
	if erro != nil {
		return erro
	}

	return nil
}
