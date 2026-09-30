напиши функцию IsValidUsername(name string) (bool, error), которая:

Возвращает false и ошибку, если имя пустое.
Возвращает false и ошибку, если длина имени меньше 3 или больше 20 символов.
ошибки возвращаются так return errors.New("message")