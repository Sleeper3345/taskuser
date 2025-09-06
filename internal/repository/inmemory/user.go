package inmemory

import (
    "errors"
    "taskuser/internal/domain/users/models"
)

type UserRepository struct {
    users  []models.User
    nextID int
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        users:  make([]models.User, 0),
        nextID: 1,
    }
}

func (r *UserRepository) GetAll() []models.User {
    return r.users
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
    for _, u := range r.users {
        if u.ID == id {
            return &u, nil
        }
    }

    return nil, errors.New("Запись не найдена")
}

func (r *UserRepository) Create(user models.User) models.User {
    user.ID = r.nextID
    r.nextID++
    r.users = append(r.users, user)

    return user
}

func (r *UserRepository) Update(id int, updated models.User) (*models.User, error) {
    for i, u := range r.users {
        if u.ID == id {
            updated.ID = id
            r.users[i] = updated
            
            return &r.users[i], nil
        }
    }

    return nil, errors.New("Запись не найдена")
}

func (r *UserRepository) Delete(id int) error {
    for i, u := range r.users {
        if u.ID == id {
            r.users = append(r.users[:i], r.users[i+1:]...)
            return nil
        }
    }

    return errors.New("Запись не найдена")
}
