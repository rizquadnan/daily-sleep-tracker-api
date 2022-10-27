CREATE TABLE sleep (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  date_of_sleep DATE NOT NULL,
  sleep_start TIME NOT NULL,
  sleep_end TIME NOT NULL,
  user_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES userperson(id) ON DELETE CASCADE 
)

INSERT INTO sleep (date_of_sleep, sleep_start, sleep_end, user_id)
VALUES (DATE '14-10-2022', '21:00', '05:00', 3),
  (DATE '15-10-2022', '20:32', '04:30', 3),
  (DATE '16-10-2022', '22:12', '05:00', 3),
  (DATE '09-10-2022', '23:21', '05:24', 4), 
  (DATE '10-10-2022', '22:55', '08:00', 5);