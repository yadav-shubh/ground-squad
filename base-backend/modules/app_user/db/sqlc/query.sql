-- name: CreateRole :one
INSERT INTO auth.app_user_roles (
    role_name, description,
    created_at, updated_at,
    created_by, updated_by
) VALUES (
             $1, $2,
             (extract(epoch from now()) * 1000)::bigint,
             (extract(epoch from now()) * 1000)::bigint,
             $3, $4
         )
RETURNING *;

-- name: GetRoleByID :one
SELECT *
FROM auth.app_user_roles
WHERE id = $1
LIMIT 1;

-- name: GetRoleByName :one
SELECT *
FROM auth.app_user_roles
WHERE role_name = $1
LIMIT 1;

-- name: ListRoles :many
SELECT *
FROM auth.app_user_roles
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateRole :one
UPDATE auth.app_user_roles
SET
    role_name = COALESCE($2, role_name),
    description = COALESCE($3, description),
    updated_by = COALESCE($4, updated_by)
WHERE id = $1
RETURNING *;

-- name: CreateAppUser :one
INSERT INTO auth.app_users (
    first_name, last_name, mobile_no, preferred_user_name,
    email, role_id, is_active, is_deleted,
    fcm_token, created_at, updated_at,
    created_by, updated_by, last_active_date
) VALUES (
             $1, $2, $3, $4,
             $5, $6, $7, $8,
             $9,
             (extract(epoch from now()) * 1000)::bigint,
             (extract(epoch from now()) * 1000)::bigint,
             $10, $11, $12
         )
RETURNING *;

-- name: GetAppUserByID :one
SELECT *
FROM auth.app_users
WHERE id = $1
LIMIT 1;

-- name: GetAppUserByEmail :one
SELECT *
FROM auth.app_users
WHERE email = $1
LIMIT 1;

-- name: GetAppUserByMobile :one
SELECT *
FROM auth.app_users
WHERE mobile_no = $1
LIMIT 1;

-- name: GetAppUserByPreferredUserName :many
SELECT *
FROM auth.app_users
WHERE preferred_user_name = $1
ORDER BY id;

-- name: ListAppUsers :many
SELECT *
FROM auth.app_users
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListUsersByRole :many
SELECT *
FROM auth.app_users
WHERE role_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: ListActiveUsers :many
SELECT *
FROM auth.app_users
WHERE is_active = TRUE
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListDeletedUsers :many
SELECT *
FROM auth.app_users
WHERE is_deleted = TRUE
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: UpdateAppUser :one
UPDATE auth.app_users
SET
    first_name = COALESCE($2, first_name),
    last_name = COALESCE($3, last_name),
    mobile_no = COALESCE($4, mobile_no),
    preferred_user_name = COALESCE($5, preferred_user_name),
    email = COALESCE($6, email),
    role_id = COALESCE($7, role_id),
    is_active = COALESCE($8, is_active),
    is_deleted = COALESCE($9, is_deleted),
    fcm_token = COALESCE($10, fcm_token),
    updated_by = COALESCE($11, updated_by),
    last_active_date = COALESCE($12, last_active_date)
WHERE id = $1
RETURNING *;
