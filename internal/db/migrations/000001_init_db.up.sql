CREATE TABLE IF NOT EXISTS peer
(
    id      uuid PRIMARY KEY NOT NULL,
    address varchar UNIQUE   NOT NULL
);

CREATE TABLE IF NOT EXISTS file
(
    hash varchar PRIMARY KEY NOT NULL
);

CREATE TABLE IF NOT EXISTS peer_file
(
    peer_id uuid REFERENCES peer (id) ON UPDATE CASCADE ON DELETE CASCADE,
    hash    varchar REFERENCES file (hash) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT peer_file_pk PRIMARY KEY (peer_id, hash)
);

CREATE INDEX peer_index ON peer_file (peer_id);

CREATE INDEX hash_index ON peer_file (hash);
