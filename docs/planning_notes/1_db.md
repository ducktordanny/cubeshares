# DataBase, Tables, etc.

Porbably needed Tables and their attributes:

- User

  - id
  - wca_id
  - name
  - email
  - avatar_url
  - role
  - created_at

- Follow

  - follower_id
  - following_id
  - created_at

- PuzzleCategory

  - id
  - name

- Solve

  - id
  - user_id
  - category_id
  - penalty_id
  - result (not calling this time, allows more flexibility for e.g. fmc)
  - scramble
  - solution
  - note

- Average

  - id
  - user_id
  - average_time
  - note

- Average Solve

  - id
  - average_id
  - solve_id

- Penalty

  - id
  - type

- Post

  - id
  - user_id
  - solve_id
  - average_id
  - description
  - created_at

- Like

  - id
  - user_id
  - post_id
  - created_at

- Comment

  - id
  - user_id
  - post_id
  - content
  - created_at

ER Diagram of database: See below or [in Lucid project](https://lucid.app/lucidchart/ab6226c4-6f45-43fa-b24a-b640c2b84a78/edit?viewport_loc=-92%2C-158%2C2380%2C1295%2C0_0&invitationId=inv_375ddb79-843f-4e56-aa14-135a594d897a)

![Image for: ER Diagram of database](../assets/database_er_diagram.svg)
