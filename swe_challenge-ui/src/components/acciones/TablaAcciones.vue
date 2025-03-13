<template>
  <div class="min-h-screen flex items-center justify-center bg-blue-50 p-6">
    <div class="overflow-x-auto w-full">
      <!-- Tabla con filtros avanzados -->
      <DataTable
        class="tabla-acciones"
        :filters="filters"
        :value="acciones"
        paginator
        stripedRows 
        :rows="5"
        :rowsPerPageOptions="[5, 10, 20, 50, 100]"
        dataKey="ticker"
        filterDisplay="menu"
        :loading="loading"
        :globalFilterFields="['ticker', 'empresa', 'corretaje', 'accion', 'rating_from', 'rating_to']"
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
        <Column field="objetivo_desde" header="Objetivo Desde" sortable filterField="objetivo_desde">
          <template #body="slotProps">
            {{ formatCurrency(slotProps.data.objetivo_desde) }}
          </template>
        </Column>
        <Column field="objetivo_a" header="Objetivo A" sortable filterField="objetivo_a">
          <template #body="slotProps">
            {{ formatCurrency(slotProps.data.objetivo_a) }}
          </template>
        </Column>
      </DataTable>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useTablaAcciones } from "./TablaAcciones";

// 🔹 Importa los componentes de PrimeVue dentro del `.vue`
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";

const { acciones, filters, loading } = useTablaAcciones();

const formatCurrency = (value: number) => {
  return new Intl.NumberFormat("es-CO", { style: "currency", currency: "COP" }).format(value);
};

</script>

<style scoped src="./TablaAcciones.css"></style>


