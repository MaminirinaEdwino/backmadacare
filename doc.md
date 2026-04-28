# Documentation API

## ambulance

**Endpoint:** `GET` `{{host}}/ambulances`  

---
## delete ambulances

**Endpoint:** `DELETE` `{{host}}/ambulances/2`  

---
## delete etablissement

**Endpoint:** `DELETE` `{{host}}/etablissements/1`  

---
## delete patient

**Endpoint:** `DELETE` `{{host}}/patients/1`  

---
## delete personnel

**Endpoint:** `DELETE` `{{host}}/personnel/4`  

---
## etablissement

**Endpoint:** `GET` `{{host}}/etablissements`  

---
## get admin

**Endpoint:** `GET` `{{host}}/admins`  

---
## patient

**Endpoint:** `GET` `{{host}}/patients`  

---
## personnel

**Endpoint:** `GET` `{{host}}/personnel`  

---
## post admin

**Endpoint:** `POST` `{{host}}/admins`  
### Request Body (JSON)
```json
{
    "nom": "admin nom",
    "prenom": "admin prenom",
    "username": "username admin",
    "mdp": "administrateur",
    "email": "admin@gmail.com",
    "etablissement_id": 2
}
```
### Headers
- **Content-Type**: application/json

---
## post ambulances

**Endpoint:** `POST` `{{host}}/ambulances`  
### Request Body (JSON)
```json
{
    "refference": "ambul 2",
    "chauffeur_id": 5,
    "status": "libre"
}
```
### Headers
- **Content-Type**: application/json

---
## post etablissement

**Endpoint:** `POST` `{{host}}/etablissements`  
### Request Body (JSON)
```json
{
    "nom": "hopitaly be 2",
    "region": "vakinakaratra",
    "contact": "034 34 344 34",
    "categorie": "CSB2"
}
```
### Headers
- **Content-Type**: application/json

---
## post patient

**Endpoint:** `POST` `{{host}}/patients`  
### Request Body (JSON)
```json
{
    "nom": "pat 12",
    "prenom": "pat 1 prenom",
    "maladies": "Grippe",
    "etablissement_id": 2,
    "status": "en_attente"
}
```
### Headers
- **Content-Type**: application/json

---
## post personnel

**Endpoint:** `POST` `{{host}}/personnel`  
### Request Body (JSON)
```json
{
    "nom": "perso chauf",
    "prenom": "perso prenom 1 chauf",
    "contact": "034 34 344 34",
    "poste": "chauffeur",
    "age": 54,
    "etablissement_id": 2
}
```
### Headers
- **Content-Type**: application/json

---
## predict

**Endpoint:** `POST` `{{host}}/predict`  
### Request Body (JSON)
```json
{
    "evidence": {
        "deshydratation": 1,
        "diarrhee_severe": 0
    },
    "target": "Maladie"
}
```
### Headers
- **Content-Type**: application/json

---
## put ambulances

**Endpoint:** `PUT` `{{host}}/ambulances/1`  
### Request Body (JSON)
```json
{
    "refference": "ambul 1",
    "chauffeur_id": 5,
    "status": "occupe"
}
```
### Headers
- **Content-Type**: application/json

---
## put etablissement

**Endpoint:** `PUT` `{{host}}/etablissements/1`  
### Request Body (JSON)
```json
{
    "nom": "hopitaly be put",
    "region": "vakinakaratra",
    "contact": "034 34 344 34",
    "categorie": "CSB2"
}
```
### Headers
- **Content-Type**: application/json

---
## put patient

**Endpoint:** `PUT` `{{host}}/patients/1`  
### Request Body (JSON)
```json
{
    "nom": "pat 12 put",
    "prenom": "pat 1 prenom",
    "maladies": "Grippe",
    "etablissement_id": 2,
    "status": "en_attente"
}
```
### Headers
- **Content-Type**: application/json

---
## put personnel

**Endpoint:** `PUT` `{{host}}/personnel/4`  
### Request Body (JSON)
```json
{
    "nom": "perso 2 put 4",
    "prenom": "perso prenom 1",
    "contact": "034 34 344 34",
    "poste": "Infirmier",
    "age": 53,
    "etablissement_id": 2
}
```
### Headers
- **Content-Type**: application/json

---
