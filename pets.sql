-- sqlite3 pets.db -init pets.sql  
BEGIN;

CREATE TABLE IF NOT EXISTS pets(
  id integer PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  animal_specie text NOT NULL,
  birth_date text NOT NULL
);

COMMIT;
