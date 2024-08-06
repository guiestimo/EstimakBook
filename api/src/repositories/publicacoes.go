package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (p Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := p.db.Prepare(
		"INSERT INTO PUBLICACOES (titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (p Publicacoes) BuscarPorId(publicacaoId uint64) (models.Publicacao, error) {
	linhas, erro := p.db.Query(`
		SELECT p.*, u.nick FROM
		publicacoes p INNER JOIN usuarios u
		ON u.id = p.autor_id WHERE p.id = ?`, publicacaoId)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao models.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (p Publicacoes) Buscar(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := p.db.Query(`
		SELECT distinct p.*, u.nick FROM PUBLICACOES p
		INNER JOIN USUARIOS u ON u.ID = p.autor_id
		INNER JOIN SEGUIDORES s ON p.autor_id = s.usuario_id
		WHERE p.autor_id = ? OR s.seguidor_id = ? 
		ORDER BY 1 DESC`, usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (p Publicacoes) Atualizar(publicacaoId uint64, publicacao models.Publicacao) error {
	statement, erro := p.db.Prepare(
		"UPDATE PUBLICACOES SET titulo = ?, conteudo = ? WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (p Publicacoes) Excluir(publicacaoId uint64) error {
	statement, erro := p.db.Prepare(
		"DELETE FROM publicacoes WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (p Publicacoes) BuscarPorUsuario(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := p.db.Query(`
	SELECT p.*, u.nick FROM PUBLICACOES p
	JOIN USUARIOS u ON u.id = p.autor_id
	WHERE p.autor_id = ?`, usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.Id,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorId,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (p Publicacoes) Curtir(publicacaoId uint64) error {
	statement, erro := p.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

func (p Publicacoes) Descurtir(publicacaoId uint64) error {
	statement, erro := p.db.Prepare(`
		UPDATE publicacoes SET curtidas = 
		CASE
			WHEN curtidas > 0 THEN curtidas - 1
			ELSE 0 END 
		WHERE id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
