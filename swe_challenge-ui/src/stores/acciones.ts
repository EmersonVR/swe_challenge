import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useAccionesStore = defineStore("acciones", () => {
  const acciones = ref<any[]>([]);
  const recomendaciones = ref<any[]>([]);

  async function fetchAcciones() {
    try {
      const response = await fetch("http://localhost:9090/acciones");
      acciones.value = await response.json();
    } catch (error) {
      console.error("Error al obtener acciones:", error);
    }
  }

  async function fetchRecomendaciones() {
    try {
      const response = await fetch("http://localhost:9090/recomendaciones");
      recomendaciones.value = await response.json();
    } catch (error) {
      console.error("Error al obtener recomendaciones:", error);
    }
  }

  // 🔥 Computed para filtrar acciones con mejores puntuaciones
  const mejoresAcciones = computed(() =>
    recomendaciones.value.sort((a, b) => b.puntuacion - a.puntuacion).slice(0, 5)
  );

  return { acciones, recomendaciones, fetchAcciones, fetchRecomendaciones, mejoresAcciones };
});
