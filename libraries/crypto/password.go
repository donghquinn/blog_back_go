package crypto

import "golang.org/x/crypto/bcrypt"

// 평문의 패스워드에서 단방향 해시를 생성한다
func EncryptHashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
}

// 패스워드 비교
func PasswordCompare(hash, password string) (bool, error) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        if err == bcrypt.ErrMismatchedHashAndPassword {
            // MEMO: err를 wrap 하여 상세를 전달하면 좋다
            return false, err
        }
        return false, err
    }
    return true, nil
}