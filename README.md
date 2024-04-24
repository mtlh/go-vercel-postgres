# Go + Vercel Functions + Postgres 

Instructions on how to create an api layer with Go, all hosted on vercel and connecting to a postgres database.

<table>
  <tr>
    <td align="center">
      <a href="https://www.postgresql.org/">
        <img src="https://cdn.clever-cloud.com/uploads/2023/08/pgsql.svg" alt="Postgres" width="140">
      </a>
    </td>
    <td align="center" style="background-color: white;">
      <a href="https://vercel.com/">
        <img src="https://encore.dev/assets/resources/vercel_cover.jpg" alt="Vercel" width="140">
      </a>
    </td>
    <td align="center" style="background-color: white;">
      <a href="https://go.dev/">
        <img src="https://i.pinimg.com/564x/9c/1a/7b/9c1a7b98ba1e02023393846c9509c587.jpg" alt="Go" width="140">
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
