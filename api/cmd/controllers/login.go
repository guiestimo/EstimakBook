package controllers

import (
	"api/cmd/auth"
	"api/cmd/banco"
	"api/cmd/models"
	"api/cmd/repositories"
	"api/cmd/responses"
	"api/cmd/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// Login
//
//		@Summary		Login
//		@Description	Login
//		@Tags			login
//		@Accept			json
//		@Produce		json
//		@Param usuario body models.Usuario true "Login Credentials"
//		@Success		200 {object} models.DadosAutenticacao
//		@Failure		422 {object} error
//		@Failure		400 {object} error
//		@Failure		401 {object} error
//	 	@Failure		500 {object} error
//		@Router			/login [get]
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(bodyRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioId := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)
	dadosAutenticacao := models.DadosAutenticacao{ID: usuarioId, Token: token}
	responses.JSON(w, http.StatusOK, dadosAutenticacao)
}
