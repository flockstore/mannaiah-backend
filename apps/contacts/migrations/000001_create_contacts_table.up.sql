CREATE TABLE contacts (
                          id TEXT PRIMARY KEY,
                          doc_type TEXT,
                          doc_number TEXT,
                          legal_name TEXT,
                          first_name TEXT,
                          last_name TEXT,
                          address TEXT,
                          address_extra TEXT,
                          city_code TEXT,
                          phone TEXT,
                          email TEXT,
                          created_at TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP NOT NULL,
                          deleted_at TIMESTAMP

);

CREATE INDEX idx_contacts_deleted_at ON contacts (deleted_at);