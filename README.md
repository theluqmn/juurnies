# Juurnies - progress > outcome

We believe life is a continuous story - a journey. In the past, sharing stories
contradicts reality. We only share select points of the story
(typically the ending/outcome), whilst reality changes continually. To combat
this problem, Juurnies is developed with the vision to popularise sharing
progress and updates rather than outcomes or success.

- Create journals on anything! (travels, projects, you name it).
- Design pages for your journals.
- Fully-free and open-source.

## Technical Overview

Project stack:

- languages: Go, HTML
- libraries: HTMX, echo, Tailwind CSS, SQLite
- server: Hack Club Nest

Database schema:

- `users(id, name, password, config`)
- `journeys(id, author, title, description, config)`
- `posts(id, journey, author, title, description, meta)`

## Notes

Developed by [Luqman](https://theluqmn.github.io). Licensed under the MIT License.
