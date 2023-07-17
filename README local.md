docker compose up -d db

docker exec -it db psql -U postgres

CREATE TABLE "akun" (
  "id" SERIAL PRIMARY KEY,
  "id_pengguna" INTEGER NOT NULL,
  "nama" VARCHAR(50) NOT NULL,
  "password" VARCHAR(100) NOT NULL,
  "tgl_daftar" DATE NOT NULL,
  "id_peran" INTEGER NOT NULL
);

CREATE TABLE "pengguna" (
  "id" SERIAL PRIMARY KEY,
  "nama" VARCHAR(50) NOT NULL,
  "alamat" VARCHAR(100) NOT NULL,
  "kode_pos" VARCHAR(255) NOT NULL,
  "provinsi" VARCHAR(50) NOT NULL,
  "kantor" INTEGER NOT NULL
);

CREATE TABLE "kantor" (
  "id" SERIAL PRIMARY KEY,
  "kode_cabang" VARCHAR(50) NOT NULL,
  "nama_cabang" VARCHAR(50) NOT NULL
);

CREATE TABLE "peran" (
  "id" SERIAL PRIMARY KEY,
  "kode_peran" VARCHAR(50) NOT NULL,
  "nama_peran" VARCHAR(50) NOT NULL
);

CREATE TABLE "layar" (
  "id" SERIAL PRIMARY KEY,
  "kode_layar" VARCHAR(50) NOT NULL,
  "nama_layar" VARCHAR(50) NOT NULL
);

CREATE TABLE "peranLayar" (
  "id" SERIAL PRIMARY KEY,
  "id_peran" INTEGER NOT NULL,
  "id_layar" INTEGER NOT NULL
);

docker compose build

docker compose up csharpapp


📝 Create a user

To create a new user, make a POST request to localhost:8080/api/users.

The body of the request should be like that:

{
    "name": "aaa",
    "email": "aaa@mail"
}

📝 Get a user

To get a user, make a GET request to localhost:8000/api/users/{id}.

For example GET request to localhost:8000/api/users/1.

📝 Update a user

To update a user, make a PUT request to localhost:8000/api/users/{id}.

For example PUT request to localhost:8000/api/users/2.

📝 Delete a user

To delete a user, make a DELETE request to localhost:8000/api/users/{id}.

For example DELETE request to localhost:8000/api/users/1.



On Postman you should see something like that:

https://dev.to/francescoxx/c-c-sharp-crud-rest-api-using-net-7-aspnet-entity-framework-postgres-docker-and-docker-compose-493a