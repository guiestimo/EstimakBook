package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requisicoes"
)

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioId, r)
	go BuscarSeguidores(canalSeguidores, usuarioId, r)
	go BuscarSeguindo(canalSeguindo, usuarioId, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioId, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar o usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregado := <-canalSeguidores:
			if seguidoresCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar os seguidores")
			}

			seguidores = seguidoresCarregado

		case seguindoCarregado := <-canalSeguindo:
			if seguindoCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar quem o usuário está seguindo")
			}

			seguindo = seguindoCarregado

		case publicacoesCarregado := <-canalPublicacoes:
			if publicacoesCarregado == nil {
				return Usuario{}, errors.New("erro ao buscar publicacoes")
			}

			publicacoes = publicacoesCarregado
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}
	canal <- usuario
}

func BuscarSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguidores
}

func BuscarSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguindo
}

func BuscarPublicacoes(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
