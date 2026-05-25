package models

import (
	"errors"
	"strings"
	"github.com/badoux/checkmail"
)

// CPFValidator é uma função simples para validar CPF
func CPFValidator(cpf string) error {
	// Verifica se CPF está vazio
	if cpf == "" {
		return errors.New("O CPF é obrigatório")
	}

	// Remove caracteres não numéricos
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verifica se tem 11 dígitos
	if len(cpf) != 11 {
		return errors.New("CPF deve ter 11 dígitos")
	}

	// Verifica se todos os dígitos são iguais (CPFs inválidos)
	if strings.Count(cpf, string(cpf[0])) == 11 {
		return errors.New("CPF inválido")
	}

	// Validação simples do primeiro dígito verificador
	// (Aqui você pode implementar a validação completa do CPF)

	return nil
}

type Users struct {
	ID       int64  `json:"id"`
	Name     string `json:"nome_usuario"`
	CPF      string `json:"cpf"`
	Email    string `json:"email_usuario"`
	Password string `json:"-"` // Não expõe a senha no JSON
}

func (u *Users) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *Users) validate(step string) error {
	if u.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	// Verifica formato do e-mail
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	// Valida CPF
	if err := CPFValidator(u.CPF); err != nil {
		return err
	}

	if step == "create" && u.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *Users) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	u.Name = strings.ToLower(u.Name)
	u.Email = strings.ToLower(u.Email)

	// Remove máscara do CPF (pontos e traços)
	u.CPF = strings.ReplaceAll(u.CPF, ".", "")
	u.CPF = strings.ReplaceAll(u.CPF, "-", "")

	return nil
}
