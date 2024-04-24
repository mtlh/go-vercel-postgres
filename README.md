# Go + Vercel Functions + Postgres 

Instructions on how to create an api layer with Go, all hosted on vercel and connecting to a postgres database.

<table>
  <tr>
    <td align="center" style="background: white;">
      <a href="https://www.postgresql.org/">
        <img src="https://static-00.iconduck.com/assets.00/postgresql-icon-1987x2048-v2fkmdaw.png" alt="Postgres" width="100">
      </a>
    </td>
    <td align="center" style="background: white;">
      <a href="https://vercel.com/">
        <img src="https://static-00.iconduck.com/assets.00/vercel-icon-512x449-3422jidz.png" alt="Vercel" width="100">
      </a>
    </td>
    <td align="center" style="background: white;">
      <a href="https://go.dev/">
        <img src="https://static-00.iconduck.com/assets.00/golang-icon-398x512-eygvdisi.png" alt="Go" width="100">
      </a>
    </td>
  </tr>
</table>

## Routes

### /api
<a href="https://go-vercel-postgres.vercel.app/api">https://go-vercel-postgres.vercel.app/api</a>
<br>
Index for application, holding basic application information.

### /api/dbtest
<a href="https://go-vercel-postgres.vercel.app/api/dbtest">https://go-vercel-postgres.vercel.app/api/dbtest</a>
<br>
Connect to the postgres instance and retrieve id and name fields in the "USERS" table.

## Getting Started

1. Clone the repo:
   ```sh
    git clone https://github.com/mtlh/go-vercel-postgres.git
   ```

2. Create .env file with the following fields:
   ```
    PGHOST='hosturl'
    PGDATABASE='db'
    PGUSER='user'
    PGPASSWORD='pass'
   ```

3. Link to Vercel:
    ```
    npm -g vercel
    vercel --prod
   ```
