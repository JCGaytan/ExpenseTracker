# Personal Expense Tracker

<!-- Language Navigator -->
<p align="right">
  <b>Languages:</b>
  <a href="#en">EN</a> |
  <a href="#es">ES</a> |
  <a href="#fr">FR</a>
</p>

---

<a id="en"></a>
## 🌍 English

A modern full-stack expense tracking solution built with Go and Vue.js, designed to help individuals and small teams manage their personal finances effectively.

## 🎯 Project Overview

This repository contains a complete expense management system consisting of two main components:

- **[GoRestAPI](./GoRestAPI/)** - A robust REST API backend built with Go
- **[Admin-Client](./admin-client/)** - A modern Vue.js 3 admin dashboard

## 🚀 Projects Structure

```
ExpenseTracker/
├── GoRestAPI/           # Backend REST API (Go + Gin + GORM)
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   ├── migrations/
│   ├── docs/
│   └── README.md       # Detailed backend documentation
├── admin-client/        # Frontend Admin Dashboard (Vue 3 + Vite)
│   ├── src/
│   ├── public/
│   └── README.md       # Detailed frontend documentation
└── README.md           # This file
```

## 🛠️ Technology Stack

### Backend (GoRestAPI)
- **Go** with **Gin** framework for high-performance HTTP handling
- **GORM** with **SQLite** for efficient data persistence  
- **JWT Authentication** for secure user sessions
- **Swagger/OpenAPI** documentation
- **Viper** for configuration management
- **Testify** for comprehensive testing

### Frontend (Admin-Client)
- **Vue 3** with Composition API and `<script setup>` syntax
- **Vite** for fast development and optimized builds
- **Tailwind CSS** for utility-first styling
- **Heroicons** for consistent iconography
- **Native Fetch API** for HTTP communications

## ✨ Key Features

- 🔐 **Secure Authentication** - JWT-based user registration and login
- 💰 **Transaction Management** - Complete CRUD operations for income/expenses
- 🏷️ **Category System** - User-specific expense categorization
- 📊 **Financial Reports** - Monthly summaries and analytics
- 🔍 **Advanced Filtering** - Search and pagination capabilities
- 📱 **Responsive Design** - Mobile-first approach with Tailwind CSS
- 🌐 **Multi-language Support** - Documentation in EN/ES/FR
- 📖 **Interactive API Docs** - Swagger UI for API exploration

## 🚦 Quick Start

### Prerequisites
- Go 1.19+ installed
- Node.js 16+ installed
- Git for version control

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/ExpenseTracker.git
cd ExpenseTracker
```

### 2. Setup Backend (GoRestAPI)
```bash
cd GoRestAPI
go mod download
cp .env.example .env  # Configure your environment variables
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`
- Swagger UI: `http://localhost:8080/docs`

### 3. Setup Frontend (Admin-Client)
```bash
cd admin-client
npm install
npm run dev
```

The admin dashboard will be available at `http://localhost:5173`

## 📚 Detailed Documentation

For comprehensive documentation, please refer to each project's README:

- **[📖 Backend Documentation](./GoRestAPI/README.md)** - Complete API documentation, architecture details, endpoints, and deployment guide
- **[📖 Frontend Documentation](./admin-client/README.md)** - Component structure, business rules, and setup instructions

## 🔧 Development Workflow

1. **Backend Development**: Work in the `GoRestAPI/` directory
   - API endpoints and business logic
   - Database models and migrations  
   - Authentication and middleware
   - Tests and documentation

2. **Frontend Development**: Work in the `admin-client/` directory
   - Vue components and UI logic
   - API integration and state management
   - Styling with Tailwind CSS
   - User experience optimization

## 🏗️ Architecture

The application follows clean architecture principles:

```
┌─────────────────┐    HTTP/REST    ┌─────────────────┐
│   Vue.js 3      │ ◄──────────────► │   Go REST API   │
│   Admin Client  │     JSON/JWT     │   (Gin + GORM)  │
│   (Frontend)    │                  │   (Backend)     │
└─────────────────┘                  └─────────────────┘
                                              │
                                              ▼
                                     ┌─────────────────┐
                                     │   SQLite DB     │
                                     │   (Storage)     │
                                     └─────────────────┘
```

## 🔐 Security Features

- **Password Hashing** - bcrypt for secure password storage
- **JWT Tokens** - Stateless authentication with expiration
- **CORS Protection** - Configured for secure cross-origin requests
- **Input Validation** - Comprehensive validation on both client and server
- **SQL Injection Prevention** - GORM ORM with prepared statements

## 🎯 Business Use Cases

- **Personal Finance Management** - Track daily expenses and income
- **Small Business Accounting** - Monitor cash flow and categorize expenses  
- **Budget Planning** - Analyze spending patterns and financial goals
- **Learning Resource** - Demonstrate modern full-stack development

## 🚀 Deployment Options

- **Development**: Local development with hot reload
- **Production**: Docker containerization ready
- **Cloud**: Azure, AWS, or Google Cloud deployment compatible
- **Database**: Easy migration from SQLite to PostgreSQL/MySQL

## 🤝 Contributing

We welcome contributions! Please see our contributing guidelines:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

- **Issues**: Report bugs or request features via [GitHub Issues](https://github.com/yourusername/ExpenseTracker/issues)
- **Documentation**: Check individual project README files for detailed guides
- **Community**: Join discussions in our [GitHub Discussions](https://github.com/yourusername/ExpenseTracker/discussions)

## 🌟 Acknowledgments

- Built with modern Go and Vue.js ecosystems
- Inspired by clean architecture principles
- Designed for real-world financial management needs

---

**Ready to start tracking your expenses? Follow the quick start guide above and check the detailed documentation in each project folder!**

---

<a id="es"></a>
## 🌍 Español

Una solución moderna de seguimiento de gastos full-stack construida con Go y Vue.js, diseñada para ayudar a individuos y pequeños equipos a gestionar sus finanzas personales de manera efectiva.

## 🎯 Descripción del Proyecto

Este repositorio contiene un sistema completo de gestión de gastos que consta de dos componentes principales:

- **[GoRestAPI](./GoRestAPI/)** - Una robusta API REST backend construida con Go
- **[Admin-Client](./admin-client/)** - Un moderno panel de administración Vue.js 3

## 🚀 Estructura de Proyectos

```
ExpenseTracker/
├── GoRestAPI/           # API REST Backend (Go + Gin + GORM)
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   ├── migrations/
│   ├── docs/
│   └── README.md       # Documentación detallada del backend
├── admin-client/        # Panel de Admin Frontend (Vue 3 + Vite)
│   ├── src/
│   ├── public/
│   └── README.md       # Documentación detallada del frontend
└── README.md           # Este archivo
```

## 🛠️ Stack Tecnológico

### Backend (GoRestAPI)
- **Go** con framework **Gin** para manejo HTTP de alto rendimiento
- **GORM** con **SQLite** para persistencia eficiente de datos  
- **Autenticación JWT** para sesiones de usuario seguras
- **Documentación Swagger/OpenAPI**
- **Viper** para gestión de configuración
- **Testify** para testing comprensivo

### Frontend (Admin-Client)
- **Vue 3** con Composition API y sintaxis `<script setup>`
- **Vite** para desarrollo rápido y builds optimizados
- **Tailwind CSS** para estilos utility-first
- **Heroicons** para iconografía consistente
- **API Fetch Nativa** para comunicaciones HTTP

## ✨ Características Principales

- 🔐 **Autenticación Segura** - Registro de usuarios y login basado en JWT
- 💰 **Gestión de Transacciones** - Operaciones CRUD completas para ingresos/gastos
- 🏷️ **Sistema de Categorías** - Categorización de gastos específica por usuario
- 📊 **Reportes Financieros** - Resúmenes mensuales y analíticas
- 🔍 **Filtrado Avanzado** - Capacidades de búsqueda y paginación
- 📱 **Diseño Responsivo** - Enfoque mobile-first con Tailwind CSS
- 🌐 **Soporte Multi-idioma** - Documentación en EN/ES/FR
- 📖 **Documentación API Interactiva** - UI Swagger para exploración de API

## 🚦 Inicio Rápido

### Prerequisitos
- Go 1.19+ instalado
- Node.js 16+ instalado
- Git para control de versiones

### 1. Clonar el Repositorio
```bash
git clone https://github.com/yourusername/ExpenseTracker.git
cd ExpenseTracker
```

### 2. Configurar Backend (GoRestAPI)
```bash
cd GoRestAPI
go mod download
cp .env.example .env  # Configurar variables de entorno
go run cmd/server/main.go
```

La API estará disponible en `http://localhost:8080`
- UI Swagger: `http://localhost:8080/docs`

### 3. Configurar Frontend (Admin-Client)
```bash
cd admin-client
npm install
npm run dev
```

El panel de administración estará disponible en `http://localhost:5173`

## 📚 Documentación Detallada

Para documentación comprensiva, consulte el README de cada proyecto:

- **[📖 Documentación Backend](./GoRestAPI/README.md)** - Documentación completa de API, detalles de arquitectura, endpoints y guía de despliegue
- **[📖 Documentación Frontend](./admin-client/README.md)** - Estructura de componentes, reglas de negocio e instrucciones de configuración

## 🔧 Flujo de Desarrollo

1. **Desarrollo Backend**: Trabajar en el directorio `GoRestAPI/`
   - Endpoints de API y lógica de negocio
   - Modelos de base de datos y migraciones
   - Autenticación y middleware
   - Pruebas y documentación

2. **Desarrollo Frontend**: Trabajar en el directorio `admin-client/`
   - Componentes Vue y lógica de UI
   - Integración de API y gestión de estado
   - Estilos con Tailwind CSS
   - Optimización de experiencia de usuario

## 🏗️ Arquitectura

La aplicación sigue principios de arquitectura limpia:

```
┌─────────────────┐    HTTP/REST    ┌─────────────────┐
│   Vue.js 3      │ ◄──────────────► │   Go REST API   │
│   Admin Client  │     JSON/JWT     │   (Gin + GORM)  │
│   (Frontend)    │                  │   (Backend)     │
└─────────────────┘                  └─────────────────┘
                                              │
                                              ▼
                                     ┌─────────────────┐
                                     │   SQLite DB     │
                                     │   (Storage)     │
                                     └─────────────────┘
```

## 🔐 Características de Seguridad

- **Hashing de Contraseñas** - bcrypt para almacenamiento seguro de contraseñas
- **Tokens JWT** - Autenticación stateless con expiración
- **Protección CORS** - Configurado para peticiones cross-origin seguras
- **Validación de Entrada** - Validación comprensiva tanto en cliente como servidor
- **Prevención de Inyección SQL** - ORM GORM con declaraciones preparadas

## 🎯 Casos de Uso de Negocio

- **Gestión Financiera Personal** - Rastrear gastos diarios e ingresos
- **Contabilidad de Pequeños Negocios** - Monitorear flujo de efectivo y categorizar gastos
- **Planificación Presupuestaria** - Analizar patrones de gasto y metas financieras
- **Recurso de Aprendizaje** - Demostrar desarrollo full-stack moderno

## 🚀 Opciones de Despliegue

- **Desarrollo**: Desarrollo local con recarga en caliente
- **Producción**: Listo para contenerización Docker
- **Nube**: Compatible con despliegue Azure, AWS o Google Cloud
- **Base de Datos**: Migración fácil de SQLite a PostgreSQL/MySQL

## 🤝 Contribuyendo

¡Damos la bienvenida a contribuciones! Por favor vea nuestras pautas de contribución:

1. Hacer fork del repositorio
2. Crear tu rama de funcionalidad (`git checkout -b feature/FuncionalidadIncreable`)
3. Hacer commit de tus cambios (`git commit -m 'Agregar alguna FuncionalidadIncreable'`)
4. Hacer push a la rama (`git push origin feature/FuncionalidadIncreable`)
5. Abrir un Pull Request

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - vea el archivo [LICENSE](LICENSE) para detalles.

## 📞 Soporte

- **Issues**: Reportar bugs o solicitar funcionalidades vía [GitHub Issues](https://github.com/yourusername/ExpenseTracker/issues)
- **Documentación**: Consultar archivos README individuales de cada proyecto para guías detalladas
- **Comunidad**: Unirse a discusiones en nuestras [GitHub Discussions](https://github.com/yourusername/ExpenseTracker/discussions)

## 🌟 Reconocimientos

- Construido con ecosistemas modernos de Go y Vue.js
- Inspirado en principios de arquitectura limpia
- Diseñado para necesidades reales de gestión financiera

---

<a id="fr"></a>
## 🌍 Français

Une solution moderne de suivi des dépenses full-stack construite avec Go et Vue.js, conçue pour aider les particuliers et petites équipes à gérer efficacement leurs finances personnelles.

## 🎯 Aperçu du Projet

Ce dépôt contient un système complet de gestion des dépenses composé de deux composants principaux :

- **[GoRestAPI](./GoRestAPI/)** - Une API REST backend robuste construite avec Go
- **[Admin-Client](./admin-client/)** - Un tableau de bord d'administration Vue.js 3 moderne

## 🚀 Structure des Projets

```
ExpenseTracker/
├── GoRestAPI/           # API REST Backend (Go + Gin + GORM)
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   ├── migrations/
│   ├── docs/
│   └── README.md       # Documentation détaillée du backend
├── admin-client/        # Tableau de Bord Admin Frontend (Vue 3 + Vite)
│   ├── src/
│   ├── public/
│   └── README.md       # Documentation détaillée du frontend
└── README.md           # Ce fichier
```

## 🛠️ Stack Technologique

### Backend (GoRestAPI)
- **Go** avec le framework **Gin** pour la gestion HTTP haute performance
- **GORM** avec **SQLite** pour une persistance de données efficace  
- **Authentification JWT** pour des sessions utilisateur sécurisées
- **Documentation Swagger/OpenAPI**
- **Viper** pour la gestion de configuration
- **Testify** pour des tests complets

### Frontend (Admin-Client)
- **Vue 3** avec Composition API et syntaxe `<script setup>`
- **Vite** pour développement rapide et builds optimisés
- **Tailwind CSS** pour styles utility-first
- **Heroicons** pour iconographie cohérente
- **API Fetch Native** pour communications HTTP

## ✨ Fonctionnalités Principales

- 🔐 **Authentification Sécurisée** - Inscription utilisateur et connexion basées sur JWT
- 💰 **Gestion des Transactions** - Opérations CRUD complètes pour revenus/dépenses
- 🏷️ **Système de Catégories** - Catégorisation des dépenses spécifique par utilisateur
- 📊 **Rapports Financiers** - Résumés mensuels et analytiques
- 🔍 **Filtrage Avancé** - Capacités de recherche et pagination
- 📱 **Design Responsive** - Approche mobile-first avec Tailwind CSS
- 🌐 **Support Multi-langues** - Documentation en EN/ES/FR
- 📖 **Documentation API Interactive** - UI Swagger pour exploration d'API

## 🚦 Démarrage Rapide

### Prérequis
- Go 1.19+ installé
- Node.js 16+ installé
- Git pour contrôle de version

### 1. Cloner le Dépôt
```bash
git clone https://github.com/yourusername/ExpenseTracker.git
cd ExpenseTracker
```

### 2. Configurer le Backend (GoRestAPI)
```bash
cd GoRestAPI
go mod download
cp .env.example .env  # Configurer les variables d'environnement
go run cmd/server/main.go
```

L'API sera disponible sur `http://localhost:8080`
- UI Swagger: `http://localhost:8080/docs`

### 3. Configurer le Frontend (Admin-Client)
```bash
cd admin-client
npm install
npm run dev
```

Le tableau de bord d'administration sera disponible sur `http://localhost:5173`

## 📚 Documentation Détaillée

Pour une documentation complète, veuillez consulter le README de chaque projet :

- **[📖 Documentation Backend](./GoRestAPI/README.md)** - Documentation complète d'API, détails d'architecture, endpoints et guide de déploiement
- **[📖 Documentation Frontend](./admin-client/README.md)** - Structure des composants, règles métier et instructions de configuration

## 🔧 Flux de Développement

1. **Développement Backend** : Travailler dans le répertoire `GoRestAPI/`
   - Endpoints d'API et logique métier
   - Modèles de base de données et migrations
   - Authentification et middleware
   - Tests et documentation

2. **Développement Frontend** : Travailler dans le répertoire `admin-client/`
   - Composants Vue et logique d'interface utilisateur
   - Intégration d'API et gestion d'état
   - Styles avec Tailwind CSS
   - Optimisation de l'expérience utilisateur

## 🏗️ Architecture

L'application suit les principes d'architecture propre :

```
┌─────────────────┐    HTTP/REST    ┌─────────────────┐
│   Vue.js 3      │ ◄──────────────► │   Go REST API   │
│   Admin Client  │     JSON/JWT     │   (Gin + GORM)  │
│   (Frontend)    │                  │   (Backend)     │
└─────────────────┘                  └─────────────────┘
                                              │
                                              ▼
                                     ┌─────────────────┐
                                     │   SQLite DB     │
                                     │   (Storage)     │
                                     └─────────────────┘
```

## 🔐 Fonctionnalités de Sécurité

- **Hachage de Mots de Passe** - bcrypt pour stockage sécurisé des mots de passe
- **Tokens JWT** - Authentification sans état avec expiration
- **Protection CORS** - Configuré pour requêtes cross-origin sécurisées
- **Validation d'Entrée** - Validation complète côté client et serveur
- **Prévention d'Injection SQL** - ORM GORM avec déclarations préparées

## 🎯 Cas d'Usage Métier

- **Gestion Financière Personnelle** - Suivre les dépenses quotidiennes et revenus
- **Comptabilité Petites Entreprises** - Surveiller les flux de trésorerie et catégoriser les dépenses
- **Planification Budgétaire** - Analyser les modèles de dépenses et objectifs financiers
- **Ressource d'Apprentissage** - Démontrer le développement full-stack moderne

## 🚀 Options de Déploiement

- **Développement** : Développement local avec rechargement à chaud
- **Production** : Prêt pour la conteneurisation Docker
- **Cloud** : Compatible avec déploiement Azure, AWS ou Google Cloud
- **Base de Données** : Migration facile de SQLite vers PostgreSQL/MySQL

## 🤝 Contribuer

Nous accueillons les contributions ! Veuillez consulter nos directives de contribution :

1. Fork le dépôt
2. Créer votre branche de fonctionnalité (`git checkout -b feature/FonctionnaliteIncroyable`)
3. Commit vos changements (`git commit -m 'Ajouter une FonctionnaliteIncroyable'`)
4. Push vers la branche (`git push origin feature/FonctionnaliteIncroyable`)
5. Ouvrir une Pull Request

## 📄 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour les détails.

## 📞 Support

- **Issues** : Signaler des bugs ou demander des fonctionnalités via [GitHub Issues](https://github.com/yourusername/ExpenseTracker/issues)
- **Documentation** : Consulter les fichiers README individuels des projets pour des guides détaillés
- **Communauté** : Rejoindre les discussions dans nos [GitHub Discussions](https://github.com/yourusername/ExpenseTracker/discussions)

## 🌟 Remerciements

- Construit avec les écosystèmes modernes Go et Vue.js
- Inspiré par les principes d'architecture propre
- Conçu pour les besoins réels de gestion financière

---

**Prêt à commencer le suivi de vos dépenses ? Suivez le guide de démarrage rapide ci-dessus et consultez la documentation détaillée dans chaque dossier de projet !**