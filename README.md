# Simulador de Programación Concurrente

Este proyecto es un simulador desarrollado en Go que utiliza técnicas de programación concurrente para gestionar recursos compartidos, simulando situaciones como el acceso a espacios de estacionamiento.

## Tabla de Contenidos

- [Requisitos](#requisitos)
- [Instalación](#instalación)
- [Ejecución](#ejecución)
- [Descripción](#descripción)
- [Contribuciones](#contribuciones)
- [Licencia](#licencia)

---

## Requisitos

Asegúrate de tener instalado:

- **Go** (versión 1.20 o superior)
- Un sistema basado en **Linux** (probado en Fedora 40)
- **Git** para clonar el repositorio y gestionar versiones

## Instalación

Sigue estos pasos para configurar el entorno de desarrollo:

1. **Instalar Go**

   Si aún no tienes Go instalado, ejecuta:

   \`\`\`bash
   sudo dnf install golang -y
   \`\`\`

   Verifica la instalación:

   \`\`\`bash
   go version
   \`\`\`

   Esto debería mostrar la versión de Go instalada.

2. **Clonar el repositorio**

   Clona este repositorio en tu máquina:

   \`\`\`bash
   git clone https://github.com/JDzul24/Simulador---Programacion-Concurrente-.git
   cd Simulador---Programacion-Concurrente-
   \`\`\`

3. **Configurar las dependencias**

   Si el proyecto usa módulos de Go, asegúrate de descargar las dependencias necesarias:

   \`\`\`bash
   go mod tidy
   \`\`\`

   Esto instalará automáticamente cualquier paquete necesario para ejecutar el simulador.

## Ejecución

Para ejecutar el simulador, sigue estos pasos:

1. Asegúrate de estar en el directorio raíz del proyecto:

   \`\`\`bash
   cd Simulador---Programacion-Concurrente-
   \`\`\`

2. Ejecuta el programa con el siguiente comando:

   \`\`\`bash
   go run main.go
   \`\`\`

3. Si todo está configurado correctamente, el simulador comenzará a ejecutarse en la terminal.

## Descripción

El simulador utiliza goroutines y canales para implementar concurrencia, gestionando la interacción entre múltiples \"vehículos\" y espacios de estacionamiento disponibles. 

- **Objetivo**: Simular el acceso concurrente a recursos limitados con un enfoque escalable y eficiente.
- **Características principales**:
  - Manejo de sincronización entre múltiples procesos.
  - Ejemplo claro de uso de goroutines y canales en Go.
  - Escenarios aleatorios para garantizar resultados variados en cada ejecución.

## Contribuciones

¡Las contribuciones son bienvenidas! Si deseas mejorar este proyecto, sigue estos pasos:

1. Haz un fork de este repositorio.
2. Crea una nueva rama para tus cambios:

   \`\`\`bash
   git checkout -b feature-nueva-funcionalidad
   \`\`\`

3. Haz tus cambios y realiza un commit:

   \`\`\`bash
   git add .
   git commit -m \"Agregar nueva funcionalidad\"
   \`\`\`

4. Sube tus cambios a tu fork:

   \`\`\`bash
   git push origin feature-nueva-funcionalidad
   \`\`\`

5. Abre un Pull Request desde tu repositorio fork al principal.

## Licencia

Este proyecto está bajo la Licencia MIT. Puedes ver más detalles en el archivo \`LICENSE\` incluido en este repositorio.

---

¡Gracias por tu interés en este proyecto! Si tienes dudas o sugerencias, no dudes en abrir un **issue** en el repositorio." > README.md
