package validate

import (
	"net/mail"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Struct нь struct-ийн validate tag-уудыг шалгана.
func Struct(val interface{}) error {
	return validate.Struct(val)
}

// IsEmail нь имэйл хаяг зөв эсэхийг шалгана.
func IsEmail(val string) bool {
	_, err := mail.ParseAddress(val)
	return err == nil
}

// IsPhoneNo нь 8 оронтой утасны дугаар эсэхийг шалгана.
func IsPhoneNo(val string) bool {
	ok, _ := regexp.MatchString(`^\d{8}$`, val)
	return ok
}

// IsRegNo нь регистрийн дугаар (2 кирилл + 8 тоо) эсэхийг шалгана.
func IsRegNo(val string) bool {
	ok, _ := regexp.MatchString(`^[а-яА-ЯөӨүҮёЁ]{2}[0-9]{8}$`, val)
	return ok
}

// IsPlateNo нь тээврийн хэрэгслийн дугаар эсэхийг шалгана.
func IsPlateNo(val string) bool {
	ok, _ := regexp.MatchString(`^[0-9]{4}[а-яА-ЯөӨүҮёЁ]{2,3}$`, val)
	return ok
}

// IsNumeric нь тоон утга эсэхийг шалгана.
func IsNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
