CREATE TABLE users (
  id SERIAL NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE TABLE feedback (
  id SERIAL NOT NULL UNIQUE,
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  category TEXT NOT NULL,
  stauts TEXT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  CONSTRAINT feedback_pk PRIMARY KEY (id),
  CONSTRAINT feedback_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE TABLE votes (
  id SERIAL NOT NULL UNIQUE,
  feedback_id INT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  CONSTRAINT votes_pk PRIMARY KEY (id),
  CONSTRAINT votes_feedback_id_fk FOREIGN KEY (feedback_id) REFERENCES feedback(id) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT votes_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE
);
CREATE TABLE comments (
  id SERIAL NOT NULL UNIQUE,
  body TEXT NOT NULL,
  feedback_id INT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  CONSTRAINT comments_pk PRIMARY KEY (id),
  CONSTRAINT comments_feedback_id_fk FOREIGN KEY (feedback_id) REFERENCES feedback(id) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT comments_user_id_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE
);