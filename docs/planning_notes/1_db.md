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
  - average_id
  - penalty_id
  - time
  - scramble
  - solution
  - note

- Average

  - id
  - user_id
  - average_time
  - note

- Penalty

  - id
  - type

- Post

  - id
  - user_id
  - solve_id
  - average_id
  - description

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

@todo Create tables and connections in Lucid
