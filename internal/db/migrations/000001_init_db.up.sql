CREATE TABLE IF NOT EXISTS peer
(
    id      uuid NOT NULL,
    address cidr NOT NULL,
    CONSTRAINT id_pk PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS file
(
    hash varchar NOT NULL,
    CONSTRAINT pk PRIMARY KEY (hash)
);

CREATE TABLE IF NOT EXISTS peer_file
(
    peer_id uuid    NOT NULL,
    hash    varchar NOT NULL,
    CONSTRAINT pk1 PRIMARY KEY (peer_id, hash),
    CONSTRAINT fk1 FOREIGN KEY (peer_id) REFERENCES peer (id),
    CONSTRAINT fk2 FOREIGN KEY (hash) REFERENCES file (hash)
);

CREATE INDEX fk3 ON peer_file (peer_id);

CREATE INDEX fk4 ON peer_file (hash);
