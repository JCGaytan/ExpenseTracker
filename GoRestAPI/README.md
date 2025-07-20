<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="en"></a>
# Personal Expense Tracker API

A backend REST API and Vue.js admin client for tracking personal incomes and expenses.

## Business Purpose
This project helps individuals or small teams track their expenses and incomes, categorize transactions, and generate summary reports. It is suitable for personal finance, family budgeting, or as a learning resource for Go REST APIs and Vue.js admin dashboards.


## Features
- User registration and login (JWT auth)
- CRUD for transactions and categories
- Filtering, pagination, and summary reports
- OpenAPI docs (Swagger UI)
- SQLite migrations
- Config via environment variables
- Admin client (Vue 3 + Vite + TailwindCSS)


## Tech Stack
- **Backend:** Go, Gin, GORM (SQLite), JWT, Viper, Swaggo
- **Frontend:** Vue 3, Vite, TailwindCSS, Heroicons
- **Testing:** Go test, Testify
- **Documentation:** Swagger/OpenAPI

---

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/JCGaytan/ExpenseTracker.git
   cd GoRestAPI
   ```

2. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   Create a `.env` file in the root directory and define the following variables:
   ```env
   DB_PATH=expense_tracker.db
   JWT_SECRET=your_secret_key
   ```

4. **Run migrations and start the server:**
   ```bash
   go run ./cmd/server
   ```

5. **Access the API documentation:**
   Open [http://localhost:8080/docs](http://localhost:8080/docs) in your browser to view Swagger UI.

6. **Run the frontend (admin-client):**
   ```bash
   cd admin-client
   npm install
   npm run dev
   ```

### Notes
- Ensure you have Go and Node.js installed on your system.
- SQLite database will be created automatically at the path specified in `DB_PATH`.
- Use a strong `JWT_SECRET` for production environments.


### Environment Variables
- `DB_PATH` (default: `expense_tracker.db`)
- `JWT_SECRET` (default: `your_secret_key`)


### Migrations
SQL files in `migrations/` are run automatically on startup. You can add new SQL files to the `migrations/` folder for schema changes.


### API Documentation
- Swagger UI available at [`/docs`](http://localhost:8080/docs)


### Testing
Run backend tests with:
```bash
go test ./...
```
Run frontend (admin-client) in dev mode:
```bash
cd admin-client
npm install
npm run dev
```

---



## Folder Structure (Monorepo)

- `cmd/server/` — Go API entrypoint. Contains the main application logic to start the server and initialize routes.
- `internal/models/` — GORM models. Defines the database schema and relationships for entities like `Transaction`, `Category`, and `User`.
- `internal/handlers/` — HTTP handlers. Implements the REST API endpoints for authentication, transactions, categories, and reports.
- `internal/middleware/` — Middleware for JWT authentication, CORS, and logging. Ensures secure and efficient request handling.
- `internal/repository/` — Database access layer. Contains functions for CRUD operations and queries using GORM.
- `internal/service/` — Business logic layer. Encapsulates application rules and processes, such as validating transactions and generating reports.
- `internal/config/` — Configuration loader, migration runner, and database seeding. Handles environment variables and initializes the application state.
- `pkg/auth/` — JWT helpers. Provides utilities for token generation, validation, and parsing.
- `migrations/` — SQL migration scripts. Contains schema changes and is automatically executed on startup.
- `docs/` — Swagger/OpenAPI documentation. Includes YAML or JSON files for API specifications.
- `admin-client/` — Vue.js admin dashboard. Implements the frontend for managing transactions and categories, using Vite and TailwindCSS.

### Considerations
- **Scalability**: The project follows clean architecture principles, separating concerns into layers (handlers, services, repositories).
- **Security**: JWT authentication and bcrypt password hashing are implemented. Ensure `JWT_SECRET` is securely stored.
- **Database**: SQLite is used for simplicity. For production, consider migrating to PostgreSQL or MySQL.
- **Frontend Integration**: CORS is enabled for seamless communication between the backend and the Vue.js admin client.
- **Testing**: Backend tests are written using Go's testing package and Testify. Frontend tests can be added using Jest or Cypress.

---



## Security
- JWT-based authentication for all protected endpoints.
- Passwords are hashed with bcrypt.
- CORS enabled for frontend integration.
- Do **not** use the default JWT secret in production.

## Business Rules
- Each user can only access their own transactions and categories.
- Amounts: positive for income, negative for expenses.
- Categories are user-specific.
- Transactions require a valid category.

## Limitations
- No multi-user admin features (each user is isolated).
- No password reset or email verification.
- SQLite is used for simplicity; not recommended for high-concurrency production.
- No mobile client (web only).

## License
MIT License. See [LICENSE](LICENSE) for details.

## Contribution
Pull requests and issues are welcome! Please open an issue to discuss major changes.

## Contact
For questions or support, open an issue or contact the maintainer.

---

## API Endpoints and Example Payloads

### Auth

#### Register
- **POST** `/auth/register`
- **Request:**
  ```json
  {
    "email": "demo@example.com",
    "password": "password123"
  }
  ```
- **Response:** `201 Created` or `400 Bad Request`

#### Login
- **POST** `/auth/login`
- **Request:**
  ```json
  {
    "email": "demo@example.com",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
    "token": "JWT_TOKEN"
  }
  ```

---

### Transactions

#### List Transactions
- **GET** `/transactions`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "amount": -50.0,
      "date": "2025-07-19T00:00:00Z",
      "category_id": 1,
      "user_id": 1,
      "description": "Demo grocery shopping"
    }
  ]
  ```

#### Get Transaction by ID
- **GET** `/transactions/{id}`
- **Response:**
  ```json
  {
    "id": 1,
    "amount": -50.0,
    "date": "2025-07-19T00:00:00Z",
    "category_id": 1,
    "user_id": 1,
    "description": "Demo grocery shopping"
  }
  ```

#### Create Transaction
- **POST** `/transactions`
- **Request:**
  ```json
  {
    "amount": -50.0,
    "date": "2025-07-19",
    "category_id": 1,
    "description": "Demo grocery shopping"
  }
  ```
- **Response:** `201 Created` with transaction object

#### Update Transaction
- **PUT** `/transactions/{id}`
- **Request:**
  ```json
  {
    "amount": -40.0,
    "date": "2025-07-20",
    "category_id": 1,
    "description": "Updated description"
  }
  ```
- **Response:** `200 OK` with updated transaction

#### Delete Transaction
- **DELETE** `/transactions/{id}`
- **Response:** `204 No Content`

---

### Categories

#### List Categories
- **GET** `/categories`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "Groceries",
      "user_id": 1
    }
  ]
  ```

#### Create Category
- **POST** `/categories`
- **Request:**
  ```json
  {
    "name": "Groceries"
  }
  ```
- **Response:** `201 Created` with category object

#### Update Category
- **PUT** `/categories/{id}`
- **Request:**
  ```json
  {
    "name": "Updated Category"
  }
  ```
- **Response:** `200 OK` with updated category

#### Delete Category
- **DELETE** `/categories/{id}`
- **Response:** `204 No Content`

---

### Reports

#### Get Monthly Summary
- **GET** `/reports/summary`
- **Response:**
  ```json
  {
    "total_income": 5000.0,
    "total_expense": -1200.0,
    "by_category": {
      "Groceries": -400.0,
      "Salary": 5000.0
    }
  }
  ```

---

You can copy and paste these endpoint and JSON blocks directly for API context or client development.

<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="es"></a>
# API de Seguimiento de Gastos Personales

Una API REST backend y cliente de administración Vue.js para el seguimiento de ingresos y gastos personales.

## Propósito del Negocio
Este proyecto ayuda a individuos o pequeños equipos a rastrear sus gastos e ingresos, categorizar transacciones y generar informes de resumen. Es adecuado para finanzas personales, presupuestos familiares o como recurso de aprendizaje para APIs REST en Go y paneles de administración en Vue.js.

## Características

- Registro y acceso de usuarios (autenticación JWT)
- CRUD para transacciones y categorías
- Filtrado, paginación e informes de resumen
- Documentación OpenAPI (Swagger UI)
- Migraciones SQLite
- Configuración por variables de entorno
- Cliente de administración (Vue 3 + Vite + TailwindCSS)

## Stack Tecnológico
- **Backend:** Go, Gin, GORM (SQLite), JWT, Viper, Swaggo
- **Frontend:** Vue 3, Vite, TailwindCSS, Heroicons
- **Testing:** Go test, Testify
- **Documentación:** Swagger/OpenAPI

---

## Comenzando

1. **Clona el repositorio:**
   ```bash
   git clone https://github.com/JCGaytan/ExpenseTracker.git
   cd GoRestAPI
   ```
2. **Instala las dependencias de Go:**
   ```bash
   go mod tidy
   ```
3. **Configura las variables de entorno:**
   Crea un archivo `.env` en la raíz y define:
   ```env
   DB_PATH=expense_tracker.db
   JWT_SECRET=your_secret_key
   ```
4. **Ejecuta migraciones e inicia el servidor:**
   ```bash
   go run ./cmd/server
   ```
5. **Accede a la documentación de la API:**
   Abre [http://localhost:8080/docs](http://localhost:8080/docs) en tu navegador.
6. **Ejecuta el frontend (admin-client):**
   ```bash
   cd admin-client
   npm install
   npm run dev
   ```

### Notas
- Asegúrate de tener Go y Node.js instalados.
- La base de datos SQLite se crea automáticamente.
- Usa un `JWT_SECRET` fuerte en producción.

### Variables de Entorno
- `DB_PATH` (por defecto: `expense_tracker.db`)
- `JWT_SECRET` (por defecto: `your_secret_key`)

### Migraciones
Los archivos SQL en `migrations/` se ejecutan automáticamente al iniciar. Puedes agregar nuevos archivos SQL para cambios de esquema.

### Documentación de la API
- Swagger UI disponible en [`/docs`](http://localhost:8080/docs)

### Testing
Ejecuta pruebas backend con:
```bash
go test ./...
```
Ejecuta el frontend en modo desarrollo:
```bash
cd admin-client
npm install
npm run dev
```

---

## Estructura de Carpetas (Monorepo)

- `cmd/server/` — Punto de entrada de la API Go. Lógica principal para iniciar el servidor y rutas.
- `internal/models/` — Modelos GORM. Define el esquema y relaciones de la base de datos.
- `internal/handlers/` — Controladores HTTP. Implementa los endpoints REST.
- `internal/middleware/` — Middleware para JWT, CORS y logging.
- `internal/repository/` — Acceso a la base de datos y operaciones CRUD.
- `internal/service/` — Lógica de negocio y validaciones.
- `internal/config/` — Carga de configuración, migraciones y seed.
- `pkg/auth/` — Utilidades JWT.
- `migrations/` — Scripts SQL de migración.
- `docs/` — Documentación Swagger/OpenAPI.
- `admin-client/` — Panel de administración Vue.js.

### Consideraciones
- **Escalabilidad**: Arquitectura limpia, separación de capas.
- **Seguridad**: JWT y bcrypt. Protege tu `JWT_SECRET`.
- **Base de datos**: SQLite por simplicidad. Considera PostgreSQL/MySQL para producción.
- **Integración Frontend**: CORS habilitado.
- **Testing**: Pruebas backend con Go y Testify. Frontend puede usar Jest o Cypress.

---

## Seguridad
- Autenticación JWT para endpoints protegidos.
- Contraseñas hasheadas con bcrypt.
- CORS habilitado.
- No uses el JWT secreto por defecto en producción.

## Reglas de Negocio
- Cada usuario accede solo a sus transacciones y categorías.
- Montos: positivos para ingresos, negativos para gastos.
- Categorías por usuario.
- Las transacciones requieren categoría válida.

## Limitaciones
- Sin funciones multiusuario admin.
- Sin recuperación de contraseña ni verificación por email.
- SQLite no recomendado para alta concurrencia.
- Solo cliente web.

## Licencia
Licencia MIT. Ver [LICENSE](LICENSE).

## Contribución
¡Pull requests y issues bienvenidos! Abre un issue para cambios mayores.

## Contacto
Para soporte, abre un issue o contacta al mantenedor.

---

## Endpoints de la API y Ejemplos

### Auth

#### Registro
- **POST** `/auth/register`
- **Request:**
  ```json
  {
    "email": "demo@example.com",
    "password": "password123"
  }
  ```
- **Response:** `201 Created` o `400 Bad Request`

#### Login
- **POST** `/auth/login`
- **Request:**
  ```json
  {
    "email": "demo@example.com",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
    "token": "JWT_TOKEN"
  }
  ```

---

### Transacciones

#### Listar Transacciones
- **GET** `/transactions`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "amount": -50.0,
      "date": "2025-07-19T00:00:00Z",
      "category_id": 1,
      "user_id": 1,
      "description": "Compra demo"
    }
  ]
  ```

#### Obtener Transacción por ID
- **GET** `/transactions/{id}`
- **Response:**
  ```json
  {
    "id": 1,
    "amount": -50.0,
    "date": "2025-07-19T00:00:00Z",
    "category_id": 1,
    "user_id": 1,
    "description": "Compra demo"
  }
  ```

#### Crear Transacción
- **POST** `/transactions`
- **Request:**
  ```json
  {
    "amount": -50.0,
    "date": "2025-07-19",
    "category_id": 1,
    "description": "Compra demo"
  }
  ```
- **Response:** `201 Created` con objeto transacción

#### Actualizar Transacción
- **PUT** `/transactions/{id}`
- **Request:**
  ```json
  {
    "amount": -40.0,
    "date": "2025-07-20",
    "category_id": 1,
    "description": "Descripción actualizada"
  }
  ```
- **Response:** `200 OK` con transacción actualizada

#### Eliminar Transacción
- **DELETE** `/transactions/{id}`
- **Response:** `204 No Content`

---

### Categorías

#### Listar Categorías
- **GET** `/categories`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "Supermercado",
      "user_id": 1
    }
  ]
  ```

#### Crear Categoría
- **POST** `/categories`
- **Request:**
  ```json
  {
    "name": "Supermercado"
  }
  ```
- **Response:** `201 Created` con objeto categoría

#### Actualizar Categoría
- **PUT** `/categories/{id}`
- **Request:**
  ```json
  {
    "name": "Categoría Actualizada"
  }
  ```
- **Response:** `200 OK` con categoría actualizada

#### Eliminar Categoría
- **DELETE** `/categories/{id}`
- **Response:** `204 No Content`

---

### Reportes

#### Obtener Resumen Mensual
- **GET** `/reports/summary`
- **Response:**
  ```json
  {
    "total_income": 5000.0,
    "total_expense": -1200.0,
    "by_category": {
      "Supermercado": -400.0,
      "Salario": 5000.0
    }
  }
  ```

---

Puedes copiar y pegar estos endpoints y bloques JSON para desarrollo de clientes o referencia de API.

<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="fr"></a>
# API de Suivi des Dépenses Personnelles

Une API REST backend et un client d'administration Vue.js pour le suivi des revenus et dépenses personnels.

## Objectif Métier
Ce projet aide les particuliers ou petites équipes à suivre leurs dépenses et revenus, catégoriser les transactions et générer des rapports de synthèse. Il convient à la gestion des finances personnelles, au budget familial ou comme ressource d'apprentissage pour les API REST Go et les tableaux de bord Vue.js.

## Fonctionnalités

- Inscription et connexion utilisateur (authentification JWT)
- CRUD pour transactions et catégories
- Filtrage, pagination et rapports de synthèse
- Documentation OpenAPI (Swagger UI)
- Migrations SQLite
- Configuration via variables d'environnement
- Client admin (Vue 3 + Vite + TailwindCSS)

## Stack Technique
- **Backend :** Go, Gin, GORM (SQLite), JWT, Viper, Swaggo
- **Frontend :** Vue 3, Vite, TailwindCSS, Heroicons
- **Tests :** Go test, Testify
- **Documentation :** Swagger/OpenAPI

---

## Démarrage

1. **Cloner le dépôt :**
   ```bash
   git clone https://github.com/JCGaytan/ExpenseTracker.git
   cd GoRestAPI
   ```
2. **Installer les dépendances Go :**
   ```bash
   go mod tidy
   ```
3. **Configurer les variables d'environnement :**
   Créez un fichier `.env` à la racine et définissez :
   ```env
   DB_PATH=expense_tracker.db
   JWT_SECRET=your_secret_key
   ```
4. **Lancer les migrations et le serveur :**
   ```bash
   go run ./cmd/server
   ```
5. **Accéder à la documentation de l'API :**
   Ouvrez [http://localhost:8080/docs](http://localhost:8080/docs) dans votre navigateur.
6. **Lancer le frontend (admin-client) :**
   ```bash
   cd admin-client
   npm install
   npm run dev
   ```

### Notes
- Assurez-vous d'avoir Go et Node.js installés.
- La base SQLite sera créée automatiquement.
- Utilisez un `JWT_SECRET` fort en production.

### Variables d'Environnement
- `DB_PATH` (défaut : `expense_tracker.db`)
- `JWT_SECRET` (défaut : `your_secret_key`)

### Migrations
Les fichiers SQL dans `migrations/` sont exécutés automatiquement au démarrage. Ajoutez de nouveaux fichiers SQL pour modifier le schéma.

### Documentation de l'API
- Swagger UI disponible sur [`/docs`](http://localhost:8080/docs)

### Tests
Lancer les tests backend :
```bash
go test ./...
```
Lancer le frontend en mode développement :
```bash
cd admin-client
npm install
npm run dev
```

---

## Structure des Dossiers (Monorepo)

- `cmd/server/` — Point d'entrée de l'API Go. Logique principale pour démarrer le serveur et les routes.
- `internal/models/` — Modèles GORM. Définit le schéma et les relations de la base de données.
- `internal/handlers/` — Handlers HTTP. Implémente les endpoints REST.
- `internal/middleware/` — Middleware pour JWT, CORS et logs.
- `internal/repository/` — Accès base de données et opérations CRUD.
- `internal/service/` — Logique métier et validations.
- `internal/config/` — Chargement config, migrations et seed.
- `pkg/auth/` — Utilitaires JWT.
- `migrations/` — Scripts SQL de migration.
- `docs/` — Documentation Swagger/OpenAPI.
- `admin-client/` — Tableau de bord Vue.js.

### Considérations
- **Scalabilité** : Architecture propre, séparation des couches.
- **Sécurité** : JWT et bcrypt. Protégez votre `JWT_SECRET`.
- **Base de données** : SQLite pour la simplicité. Privilégiez PostgreSQL/MySQL en production.
- **Intégration Frontend** : CORS activé.
- **Tests** : Backend avec Go et Testify. Frontend possible avec Jest ou Cypress.

---

## Sécurité
- Authentification JWT pour tous les endpoints protégés.
- Mots de passe hashés avec bcrypt.
- CORS activé.
- Ne pas utiliser le JWT secret par défaut en production.

## Règles Métier
- Chaque utilisateur accède uniquement à ses transactions et catégories.
- Montants : positifs pour revenus, négatifs pour dépenses.
- Catégories propres à chaque utilisateur.
- Les transactions nécessitent une catégorie valide.

## Limitations
- Pas de fonctions admin multi-utilisateur.
- Pas de réinitialisation de mot de passe ni de vérification email.
- SQLite non recommandé pour forte concurrence.
- Client web uniquement.

## Licence
Licence MIT. Voir [LICENSE](LICENSE).

## Contribution
Pull requests et issues bienvenus ! Ouvrez un issue pour les changements majeurs.

## Contact
Pour toute question, ouvrez un issue ou contactez le mainteneur.

---

## Endpoints de l'API et Exemples

### Auth

#### Inscription
- **POST** `/auth/register`
- **Request :**
  ```json
  {
    "email": "demo@example.com",
    "password": "password123"
  }
  ```
- **Response :** `201 Created` ou `400 Bad Request`

#### Connexion
- **POST** `/auth/login`
- **Request :**
  ```json
  {
    "email": "demo@example.com",
    "password": "password123"
  }
  ```
- **Response :**
  ```json
  {
    "token": "JWT_TOKEN"
  }
  ```

---

### Transactions

#### Lister les Transactions
- **GET** `/transactions`
- **Response :**
  ```json
  [
    {
      "id": 1,
      "amount": -50.0,
      "date": "2025-07-19T00:00:00Z",
      "category_id": 1,
      "user_id": 1,
      "description": "Courses demo"
    }
  ]
  ```

#### Obtenir une Transaction par ID
- **GET** `/transactions/{id}`
- **Response :**
  ```json
  {
    "id": 1,
    "amount": -50.0,
    "date": "2025-07-19T00:00:00Z",
    "category_id": 1,
    "user_id": 1,
    "description": "Courses demo"
  }
  ```

#### Créer une Transaction
- **POST** `/transactions`
- **Request :**
  ```json
  {
    "amount": -50.0,
    "date": "2025-07-19",
    "category_id": 1,
    "description": "Courses demo"
  }
  ```
- **Response :** `201 Created` avec l'objet transaction

#### Mettre à Jour une Transaction
- **PUT** `/transactions/{id}`
- **Request :**
  ```json
  {
    "amount": -40.0,
    "date": "2025-07-20",
    "category_id": 1,
    "description": "Description mise à jour"
  }
  ```
- **Response :** `200 OK` avec la transaction mise à jour

#### Supprimer une Transaction
- **DELETE** `/transactions/{id}`
- **Response :** `204 No Content`

---

### Catégories

#### Lister les Catégories
- **GET** `/categories`
- **Response :**
  ```json
  [
    {
      "id": 1,
      "name": "Courses",
      "user_id": 1
    }
  ]
  ```

#### Créer une Catégorie
- **POST** `/categories`
- **Request :**
  ```json
  {
    "name": "Courses"
  }
  ```
- **Response :** `201 Created` avec l'objet catégorie

#### Mettre à Jour une Catégorie
- **PUT** `/categories/{id}`
- **Request :**
  ```json
  {
    "name": "Catégorie Mise à Jour"
  }
  ```
- **Response :** `200 OK` avec la catégorie mise à jour

#### Supprimer une Catégorie
- **DELETE** `/categories/{id}`
- **Response :** `204 No Content`

---

### Rapports

#### Obtenir le Résumé Mensuel
- **GET** `/reports/summary`
- **Response :**
  ```json
  {
    "total_income": 5000.0,
    "total_expense": -1200.0,
    "by_category": {
      "Courses": -400.0,
      "Salaire": 5000.0
    }
  }
  ```

---

Vous pouvez copier-coller ces endpoints et blocs JSON pour le développement client ou la référence API.
