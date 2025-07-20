# Expense Tracker Admin Client

<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="en"></a>



## Business Rules

- Only authenticated users can access the admin portal. Authentication is managed via JWT tokens stored in localStorage.
- Users must log in to access the admin features. Logging off removes the token and redirects to the login page.
- The admin portal allows CRUD operations for Transactions and Categories.
- Reports are available for viewing summarized expense data.
- All navigation is handled within the admin portal sidebar.
- The logoff button is always available at the bottom of the sidebar.

## Technical Information

- **Framework:** Vue 3 (Composition API, `<script setup>` syntax)
- **Build Tool:** Vite
- **Component Structure:**
  - `App.vue`: Root component, handles authentication state and routing between login and admin portal.
  - `Login.vue`: Handles user authentication and token storage.
  - `AdminPortal.vue`: Main admin interface with sidebar navigation and logoff functionality.
  - `TransactionsCrud.vue`, `CategoriesCrud.vue`: CRUD interfaces for transactions and categories.
  - `ReportsView.vue`: Displays reports and analytics.
- **Styling:** Tailwind CSS utility classes (assumed from class names)
- **Icons:** Heroicons (imported as Vue components)
- **State Management:** Local component state with Vue's `ref` and `computed` (no Vuex/Pinia)
- **Authentication:**
  - Login form posts credentials to `/auth/login` endpoint.
  - On success, JWT token is stored in `localStorage`.
  - App checks for token to determine if user is logged in.
  - Logoff removes token and reloads the app.
- **Routing:** No Vue Router; navigation is handled by switching components in state.
- **API Communication:** Uses native `fetch` for HTTP requests.

## Project Structure

```
admin-client/
├── index.html
├── package.json
├── vite.config.js
├── public/
│   └── vite.svg
└── src/
    ├── App.vue
    ├── main.js
    ├── style.css
    ├── assets/
    │   └── vue.svg
    └── components/
        ├── AdminPortal.vue
        ├── CategoriesCrud.vue
        ├── Login.vue
        ├── ReportsView.vue
        └── TransactionsCrud.vue
```

## How to Run

1. Install dependencies:
   ```sh
   npm install
   ```
2. Start the development server:
   ```sh
   npm run dev
   ```
3. Open [http://localhost:5173](http://localhost:5173) in your browser.


- This project is the admin client for an expense tracker system. It is intended to be used by authorized users only.
- The backend API (not included here) must provide `/auth/login` and endpoints for transactions, categories, and reports.
- For production, ensure secure handling of tokens and consider using HTTPS.

**[Siguiente idioma / Next language: Español](#es-español-1)**

---


<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="es"></a>

## [ES] Español

**Referencia de idioma: Español (ES)**

### Reglas de Negocio

- Solo los usuarios autenticados pueden acceder al portal de administración. La autenticación se gestiona mediante tokens JWT almacenados en localStorage.
- Los usuarios deben iniciar sesión para acceder a las funciones de administración. Cerrar sesión elimina el token y redirige a la página de inicio de sesión.
- El portal de administración permite operaciones CRUD para Transacciones y Categorías.
- Los informes están disponibles para ver datos de gastos resumidos.
- Toda la navegación se maneja dentro de la barra lateral del portal de administración.
- El botón de cierre de sesión siempre está disponible en la parte inferior de la barra lateral.

### Información Técnica

- **Framework:** Vue 3 (API de composición, sintaxis `<script setup>`)
- **Herramienta de construcción:** Vite
- **Estructura de componentes:**
  - `App.vue`: Componente raíz, gestiona el estado de autenticación y el enrutamiento entre el inicio de sesión y el portal de administración.
  - `Login.vue`: Gestiona la autenticación del usuario y el almacenamiento del token.
  - `AdminPortal.vue`: Interfaz principal de administración con navegación lateral y funcionalidad de cierre de sesión.
  - `TransactionsCrud.vue`, `CategoriesCrud.vue`: Interfaces CRUD para transacciones y categorías.
  - `ReportsView.vue`: Muestra informes y análisis.
- **Estilos:** Clases utilitarias de Tailwind CSS (según los nombres de clase)
- **Iconos:** Heroicons (importados como componentes de Vue)
- **Gestión de estado:** Estado local del componente con `ref` y `computed` de Vue (sin Vuex/Pinia)
- **Autenticación:**
  - El formulario de inicio de sesión envía credenciales al endpoint `/auth/login`.
  - En caso de éxito, el token JWT se almacena en `localStorage`.
  - La aplicación verifica el token para determinar si el usuario ha iniciado sesión.
  - Cerrar sesión elimina el token y recarga la aplicación.
- **Enrutamiento:** Sin Vue Router; la navegación se maneja cambiando componentes en el estado.
- **Comunicación API:** Usa `fetch` nativo para solicitudes HTTP.

### Estructura del Proyecto

```
admin-client/
├── index.html
├── package.json
├── vite.config.js
├── public/
│   └── vite.svg
└── src/
    ├── App.vue
    ├── main.js
    ├── style.css
    ├── assets/
    │   └── vue.svg
    └── components/
        ├── AdminPortal.vue
        ├── CategoriesCrud.vue
        ├── Login.vue
        ├── ReportsView.vue
        └── TransactionsCrud.vue
```

### Cómo Ejecutar

1. Instalar dependencias:
   ```sh
   npm install
   ```
2. Iniciar el servidor de desarrollo:
   ```sh
   npm run dev
   ```
3. Abrir [http://localhost:5173](http://localhost:5173) en el navegador.


- Este proyecto es el cliente de administración para un sistema de seguimiento de gastos. Está destinado solo para usuarios autorizados.
- La API backend (no incluida aquí) debe proporcionar `/auth/login` y endpoints para transacciones, categorías e informes.
- Para producción, asegúrese de manejar los tokens de forma segura y considere usar HTTPS.

**[Idioma siguiente / Next language: Français](#fr-français-1)**

---


<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="fr"></a>

## [FR] Français

**Référence de langue : Français (FR)**

### Règles Métier

- Seuls les utilisateurs authentifiés peuvent accéder au portail d'administration. L'authentification est gérée via des tokens JWT stockés dans le localStorage.
- Les utilisateurs doivent se connecter pour accéder aux fonctionnalités d'administration. La déconnexion supprime le token et redirige vers la page de connexion.
- Le portail d'administration permet des opérations CRUD pour les Transactions et les Catégories.
- Des rapports sont disponibles pour consulter les données de dépenses résumées.
- Toute la navigation est gérée dans la barre latérale du portail d'administration.
- Le bouton de déconnexion est toujours disponible en bas de la barre latérale.

### Informations Techniques

- **Framework :** Vue 3 (API de composition, syntaxe `<script setup>`)
- **Outil de build :** Vite
- **Structure des composants :**
  - `App.vue` : Composant racine, gère l'état d'authentification et la navigation entre la connexion et le portail d'administration.
  - `Login.vue` : Gère l'authentification de l'utilisateur et le stockage du token.
  - `AdminPortal.vue` : Interface principale d'administration avec navigation latérale et fonctionnalité de déconnexion.
  - `TransactionsCrud.vue`, `CategoriesCrud.vue` : Interfaces CRUD pour les transactions et les catégories.
  - `ReportsView.vue` : Affiche les rapports et les analyses.
- **Style :** Classes utilitaires Tailwind CSS (d'après les noms de classes)
- **Icônes :** Heroicons (importés comme composants Vue)
- **Gestion d'état :** État local du composant avec `ref` et `computed` de Vue (sans Vuex/Pinia)
- **Authentification :**
  - Le formulaire de connexion envoie les identifiants à l'endpoint `/auth/login`.
  - En cas de succès, le token JWT est stocké dans le `localStorage`.
  - L'application vérifie le token pour déterminer si l'utilisateur est connecté.
  - La déconnexion supprime le token et recharge l'application.
- **Routage :** Pas de Vue Router ; la navigation est gérée en changeant les composants dans l'état.
- **Communication API :** Utilise `fetch` natif pour les requêtes HTTP.

### Structure du Projet

```
admin-client/
├── index.html
├── package.json
├── vite.config.js
├── public/
│   └── vite.svg
└── src/
    ├── App.vue
    ├── main.js
    ├── style.css
    ├── assets/
    │   └── vue.svg
    └── components/
        ├── AdminPortal.vue
        ├── CategoriesCrud.vue
        ├── Login.vue
        ├── ReportsView.vue
        └── TransactionsCrud.vue
```

### Comment Exécuter

1. Installer les dépendances :
   ```sh
   npm install
   ```
2. Démarrer le serveur de développement :
   ```sh
   npm run dev
   ```
3. Ouvrir [http://localhost:5173](http://localhost:5173) dans votre navigateur.


- Ce projet est le client d'administration d'un système de suivi des dépenses. Il est destiné uniquement aux utilisateurs autorisés.
- L'API backend (non incluse ici) doit fournir `/auth/login` et des endpoints pour les transactions, les catégories et les rapports.
- Pour la production, assurez-vous de gérer les tokens de manière sécurisée et envisagez d'utiliser HTTPS.

**[Next language / Langue suivante: English](#expense-tracker-admin-client)**
