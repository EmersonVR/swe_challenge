<template>
  <div class="h-screen flex flex-col items-center bg-gray-100 overflow-hidden">
    <div class="bg-white shadow-lg rounded-lg w-full h-full flex flex-col">
      <!-- Navbar de pestañas -->
      <TabView class="flex-grow">
        <!-- Pestaña de Tabla de Acciones -->
        <TabPanel header="📊 Acciones" class="h-full">
          <div class="h-full overflow-auto">
            <Suspense>
              <template #default>
                <LazyTablaAcciones />
              </template>
              <template #fallback>
                <p class="text-center text-gray-500">Cargando Acciones...</p>
              </template>
            </Suspense>
          </div>
        </TabPanel>

        <!-- Pestaña de Acciones Recomendadas -->
        <TabPanel header="🌟 Recomendaciones" class="h-full">
          <div class="h-full overflow-auto">
            <Suspense>
              <template #default>
                <LazyAccionesRecomendadas />
              </template>
              <template #fallback>
                <p class="text-center text-gray-500">Cargando Recomendaciones...</p>
              </template>
            </Suspense>
          </div>
        </TabPanel>
      </TabView>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineAsyncComponent } from "vue";
import TabView from "primevue/tabview";
import TabPanel from "primevue/tabpanel";

// Importación asíncrona de los componentes
const LazyTablaAcciones = defineAsyncComponent(() => import("./components/acciones/TablaAcciones.vue"));
const LazyAccionesRecomendadas = defineAsyncComponent(() => import("./components/accionesRecomendadas/AccionesRecomendadas.vue"));
</script>

<style >
/* Ajusta la altura de toda la página */
html, body {
   margin:0px;  
   height:100%;
   overflow: hidden;

}

/* Hace que el contenedor principal sea flexible y se ajuste */
.flex-grow {
  flex-grow: 1;
}

/* Permite que el contenido dentro de cada tab sea desplazable sin afectar la página */
.overflow-auto {
  overflow: auto;
  max-height: 100vh; /* Ajusta la altura de la tabla sin romper el diseño */
}
::-webkit-scrollbar {
    display: none;
}
/* Asegura que la app use todo el alto sin generar scroll innecesario */
.h-screen {
  height: 100vh;
  overflow: hidden;
}
</style>
