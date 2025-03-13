<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 p-6">
    <div class="bg-white rounded-lg shadow-lg w-full">

      <DataTable
        class="tabla-recomendaciones"
        :filters="filters"
        :value="recomendaciones"
        paginator
        stripedRows 
        :rows="5"
        :rowsPerPageOptions="[5, 10, 20, 50, 100]"
        dataKey="ticker"
        filterDisplay="menu"
        :loading="loading"
        :globalFilterFields="['ticker', 'empresa', 'corretaje', 'rating_from', 'rating_to', 'hora']"
      >
        <template #header>
          <div class="flex justify-between items-center">
            <div class="flex items-center gap-2">
              <i class="pi pi-search"></i>
              <InputText v-model="filters.global.value" placeholder="Buscar..." />
            </div>
          </div>
        </template>

        <Column field="ticker" header="Ticker" sortable filterField="ticker"/>
        <Column field="empresa" header="Empresa" sortable filterField="empresa"/>
        <Column field="corretaje" header="Corretaje" sortable filterField="corretaje"/>
        <Column field="accion" header="Acción" sortable filterField="accion"/>
        <Column field="rating_from" header="Rating From" sortable filterField="rating_from"/>
        <Column field="rating_to" header="Rating To" sortable filterField="rating_to"/>
        <Column field="hora" header="Hora" sortable filterField="hora">
          <template #body="{ data }">
            {{ formatFecha(data.hora) }}
          </template>
        </Column>


        <Column field="puntuacion" header="Puntuación" sortable>
          <template #body="slotProps">
            <span :class="getPuntuacionClass(slotProps.data.puntuacion)">
              {{ slotProps.data.puntuacion.toFixed(2) }}
            </span>
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAccionesRecomendadas } from "./AccionesRecomendadas";

// 🔹 Importa los componentes de PrimeVue dentro del `.vue`
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import Calendar from "primevue/calendar";

const {
  recomendaciones,
  filters,
  loading,
  getPuntuacionClass,
} = useAccionesRecomendadas();

const formatFecha = (fechaStr: string): string => {
  if (!fechaStr) return "Sin fecha";

  const fecha = new Date(fechaStr);
  return new Intl.DateTimeFormat("es-ES", {
    day: "2-digit",
    month: "short", // "feb."
    hour: "2-digit",
    minute: "2-digit",
  }).format(fecha);
};
</script>

<style scoped src="./AccionesRecomendadas.css"></style>

