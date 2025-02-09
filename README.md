<!-- PROJECT LOGO -->
<br />
div align="center">
  <h3 align="center">Event Ticketing System</h3>
  <h5 align="center">
    <br />
    <a href="https://event-ticket.me/" target="_blank">View Live</a>
    |
    <a href="https://github.com/nickemma/event-ticketing-system/issues/new" target="_blank">Report Bug</a>
    |
    <a href="https://github.com/nickemma/event-ticketing-system/issues/new" target="_blank">Request Feature</a>
  </h5>

<!-- ABOUT THE PROJECT -->

### Event Ticketing System

A scalable backend service for managing events, tickets, and transactions. Built with Golang, PostgreSQL, and Kubernetes  on the backend, and a modern Next.js frontend for a seamless user experience. Features include JWT authentication, atomic ticket inventory management, payment processing, and audit logging.

<div align="center">
  <img  width="1000" alt="screenshot" src="./images/event-ticketing-system.png">
</div>

#### Built With

**Backend & DevOps:**

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=for-the-badge&logo=google&logoColor=white)](https://grpc.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![AWS](https://img.shields.io/badge/AWS-232F3E?style=for-the-badge&logo=amazonaws&logoColor=white)](https://aws.amazon.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)](https://kubernetes.io/)
[![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)](https://github.com/features/actions)

**Frontend & UI/UX:**

[![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=next.js&logoColor=white)](https://nextjs.org/)
[![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)](https://react.dev/)
[![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwind-css&logoColor=white)](https://tailwindcss.com/)
[![Zustand](https://img.shields.io/badge/Zustand-2d2d2d?style=for-the-badge&logo=react&logoColor=white)](https://zustand-demo.pmnd.rs/)
[![React Query](https://img.shields.io/badge/React%20Query-FF4154?style=for-the-badge&logo=react-query&logoColor=white)](https://tanstack.com/query/latest)
[![Axios](https://img.shields.io/badge/Axios-5A29E4?style=for-the-badge&logo=axios&logoColor=white)](https://axios-http.com/)  
[![shadcn/ui](https://img.shields.io/badge/ShadCN_UI-000000?style=for-the-badge&logo=shadcn&logoColor=white)](https://ui.shadcn.com/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<!-- GETTING STARTED -->

### Getting Started

#### Prerequisites

**Backend**:
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL 15
- Redis 7

**Frontend**:
- Node.js 20+
- npm/yarn

#### Installation

1. **Clone the repository**
   ```sh
   git clone https://github.com/nickemma/event-ticketing-system.git
   cd event-ticketing-system

2. **Set up environment variables**
   ```sh
   cp .env.example .env

3. **Start dependencies**
   ```sh
   docker-compose up -d postgres redis

4. **Run database migrations**
   ```sh
   make migrateup

5. **Start the server**
   ```sh
   make server
   
6. **Access the App**
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- FRONTEND FEATURES -->
### Frontend Features
1. User Authentication
 - Login/register with JWT token management.
 - Protected routes for organizers.
2. Event Dashboard
- Create/edit events (organizers).
- Browse events with filters (users).
3. Ticket Purchase Flow
- Real-time inventory updates during checkout.
- Payment simulation with Stripe.js.
4. Responsive Design
- Mobile-first UI built with Tailwind CSS.

<!-- CONTRIBUTING -->

### Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/your-feature`)
3. Commit your Changes (`git commit -m 'Add some Amazing Feature'`)
4. Push to the Branch (`git push origin feature/your-feature`)
5. Open a Pull Request

Areas to Contribute:

- Improve UI/UX with animations and better styling
- Add dark mode support
- Implement more advanced authentication features
- Integrate real-time ticket availability tracking
- Add payment gateway integrations (Stripe/PayPal)
- Improve concurrent ticket purchase handling
- Enhance Kubernetes deployment manifests

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<!-- LICENSE -->

### License

Distributed under the Apache2.0 License. See [Apache License 2.0](LICENSE) for details.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<!-- CONTACT -->

### Contact

#### Nicholas Emmanuel

 <div align="center">
 <a href="https://www.linkedin.com/in/techieemma/"><img src="https://img.shields.io/badge/linkedin-%23f78a38.svg?style=for-the-badge&logo=linkedin&logoColor=white" alt="Linkedin"></a> 
 <a href="https://twitter.com/techieEmma"><img src="https://img.shields.io/badge/Twitter-%23f78a38.svg?style=for-the-badge&logo=Twitter&logoColor=white" alt="Twitter"></a> 
 <a href="https://github.com/nickemma/"><img src="https://img.shields.io/badge/github-%23f78a38.svg?style=for-the-badge&logo=github&logoColor=white" alt="Github"></a>
 <a href="mailto:nicholasemmanuel321@gmail.com"><img src="https://img.shields.io/badge/Gmail-f78a38?style=for-the-badge&logo=gmail&logoColor=white" alt="Linkedin"></a>
 </div>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

---

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments
- Inspired by [Tech School](https://github.com/techschool) for Golang backend patterns
- [Gin Web Framework](https://gin-gonic.com/) for REST API development
- [](https://sqlc.dev/) for type-safe SQL queries
- [Dribbble](https://dribbble.com/) for the UI design inspiration for this project.

<p align="right">(<a href="#readme-top">back to top</a>)</p>