#!/bin/sh
curl 'https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyAkmD0CtpwxgkHwpt6J1-_MIXne-YjL8Ew' \
-H 'Content-Type: application/json' \
--data-binary '{"email":"erdemerkoc@gmail.com","password":"123456","returnSecureToken":true}'