package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := u.db.Prepare(
		"INSERT INTO USUARIOS (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (u usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := u.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick,
		nomeOuNick)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u usuarios) BuscarPorId(id uint64) (models.Usuario, error) {
	linhas, erro := u.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", id)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (u usuarios) Atualizar(id uint64, usuario models.Usuario) error {
	statement, erro := u.db.Prepare(
		"UPDATE USUARIOS SET nome = ?, nick = ?, email = ? WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (u usuarios) Excluir(id uint64) error {
	statement, erro := u.db.Prepare(
		"DELETE FROM USUARIOS WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (u usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := u.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (u usuarios) Seguir(usuarioId uint64, seguidorId uint64) error {
	statement, erro := u.db.Prepare(
		"INSERT IGNORE INTO SEGUIDORES (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (u usuarios) PararDeSeguirUsuario(usuarioId uint64, seguidorId uint64) error {
	statement, erro := u.db.Prepare(
		"DELETE FROM SEGUIDORES WHERE usuario_id = ? AND seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (u usuarios) BuscarSeguidores(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := u.db.Query(`
	SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
	FROM seguidores S INNER JOIN usuarios u ON s.seguidor_id = u.id
	WHERE s.usuario_id = ?`, usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	if linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u usuarios) BuscarSeguindo(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := u.db.Query(`
	SELECT u.id, u.nome, u.nick, u.email, u.criadoEm
	FROM usuarios u INNER JOIN seguidores s ON u.id = s.usuario_id
	WHERE s.seguidor_id = ?`, usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	if linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := u.db.Query("SELECT SENHA FROM USUARIOS WHERE ID = ?", usuarioId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

func (u usuarios) AtualizarSenha(usuarioId uint64, senhaComHash string) error {
	statement, erro := u.db.Prepare(
		"UPDATE USUARIOS SET senha = ? WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senhaComHash, usuarioId); erro != nil {
		return erro
	}

	return nil
}
