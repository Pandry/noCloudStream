Auther needs to authenticate users via HTTP from the frontend and provide them a JWT session.  
Is also responsable of renewing the JWT session.

TODO:
- COnferma mail con expiry e confirmation token
- Session duration
- Switch per impostare iscrizioni
- Come passare il token JWT (idealmente JSON alla login e poi nelle richieste via authentication token bearer)

// There are stuff I want in the JWT, stuff I want in the DB and stuff I want the user to POST to me
// Use 3 structs

req ajax con JTW come Authorization header (risolve XSRF)

IP address in JWT?


login: JSON mail+password -> ok, jwt, error
signup: JSON name+mail+password -> ok, error


prometheus
