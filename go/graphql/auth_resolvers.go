package graphql

import (
    "context"
    "errors"
    "strconv"

    "last-go/internal/auth"
    "last-go/internal/database"

    "golang.org/x/crypto/bcrypt"
)

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
    dbUser, err := m.DB.CreateUser(input.Email, input.Username, input.Password)
    if err != nil {
        return nil, err
    }

    token, err := auth.GenerateToken(strconv.Itoa(dbUser.ID), dbUser.Email)
    if err != nil {
        return nil, err
    }

    return &AuthResponse{
        User:  mapDBUserToGQL(dbUser),
        Token: token,
    }, nil
}

func (m *mutationResolver) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
    dbUser, err := m.DB.GetUserByEmail(input.Email)
    if err != nil {
        return nil, errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)); err != nil {
        return nil, errors.New("invalid credentials")
    }

    token, err := auth.GenerateToken(strconv.Itoa(dbUser.ID), dbUser.Email)
    if err != nil {
        return nil, err
    }

    return &AuthResponse{
        User:  mapDBUserToGQL(dbUser),
        Token: token,
    }, nil
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
    uidVal := ctx.Value("userID")
    if uidVal == nil {
        return nil, nil // not authenticated
    }
    uidStr, ok := uidVal.(string)
    if !ok {
        return nil, errors.New("invalid user id in context")
    }
    uid, err := strconv.Atoi(uidStr)
    if err != nil {
        return nil, err
    }

    dbUser, err := q.DB.GetUserByID(uid)
    if err != nil {
        return nil, err
    }
    return mapDBUserToGQL(dbUser), nil
}

func (q *queryResolver) Users(ctx context.Context) ([]*User, error) {
    dbUsers, err := q.DB.GetAllUsers()
    if err != nil {
        return nil, err
    }
    out := make([]*User, 0, len(dbUsers))
    for _, u := range dbUsers {
        out = append(out, mapDBUserToGQL(u))
    }
    return out, nil
}

func mapDBUserToGQL(u *database.User) *User {
    if u == nil {
        return nil
    }
    return &User{
        ID:        strconv.Itoa(u.ID),
        Email:     u.Email,
        Username:  u.Username,
        CreatedAt: u.CreatedAt,
        UpdatedAt: u.UpdatedAt,
    }
}