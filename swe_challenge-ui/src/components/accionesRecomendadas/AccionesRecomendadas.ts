import { ref, onMounted } from "vue";
import { FilterMatchMode } from "primevue/api";
import { useAccionesStore } from "../../stores/acciones";
import dayjs from "dayjs";
import "dayjs/locale/es";

// Importar componentes de PrimeVue
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";
import Dropdown from "primevue/dropdown";
import Calendar from "primevue/calendar";

// 📌 Definimos una interfaz para el tipo de datos de recomendaciones
interface Recomendacion {
  ticker: string;
  empresa: string;
  corretaje: string;
  accion: string;
  rating_from: string;
  rating_to: string;
  hora: string; // Asumimos que es un string con formato de fecha/hora
  puntuacion: number;
}

// 📌 Tipamos la función correctamente
export function useAccionesRecomendadas() {
  const store = useAccionesStore();
  const recomendaciones = ref<Recomendacion[]>([]); // Ahora el array tiene un tipo definido
  const loading = ref<boolean>(true);

  // 📌 Definimos los filtros con `Record<string, any>` para mayor flexibilidad
  const filters = ref<Record<string, any>>({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    ticker: { value: null, matchMode: FilterMatchMode.CONTAINS },
    empresa: { value: null, matchMode: FilterMatchMode.CONTAINS },
    corretaje: { value: null, matchMode: FilterMatchMode.CONTAINS },
    accion: { value: null, matchMode: FilterMatchMode.CONTAINS },
    rating_from: { value: null, matchMode: FilterMatchMode.EQUALS },
    rating_to: { value: null, matchMode: FilterMatchMode.EQUALS },
    hora: { value: null, matchMode: FilterMatchMode.DATE_IS },
  });

  onMounted(async () => {
    await store.fetchRecomendaciones();
    recomendaciones.value = store.recomendaciones as Recomendacion[];
    loading.value = false;
  });

  // 📌 Tipamos correctamente la función que da color según la puntuación
  const getPuntuacionClass = (puntuacion: number): string => {
    if (puntuacion >= 50) return "text-green-600 font-bold";
    if (puntuacion >= 40) return "text-yellow-500 font-bold";
    return "text-red-500 font-bold";
  };

  return { recomendaciones, filters, loading, getPuntuacionClass };
}
