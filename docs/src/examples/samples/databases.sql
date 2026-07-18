CREATE TABLE roles (
  tenant_id bigint NOT NULL,
  id bigint NOT NULL,
  name varchar(100) UNIQUE,
  PRIMARY KEY (tenant_id, id)
);

CREATE TABLE users (
  id bigint PRIMARY KEY,
  tenant_id bigint NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  role_id bigint NOT NULL,
  CONSTRAINT fk_users_role FOREIGN KEY (tenant_id, role_id)
    REFERENCES roles (tenant_id, id) ON DELETE RESTRICT
);
