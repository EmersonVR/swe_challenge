import { ref, onMounted } from "vue";
import { FilterMatchMode } from "primevue/api";
import { useAccionesStore } from "../../stores/acciones";

// Importación de los componentes de PrimeVue
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import InputText from "primevue/inputtext";

// Definimos una interfaz con los datos esperados de cada acción
interface Accion {
  ticker: string;
  empresa: string;
  corretaje: string;
  accion: string;
  rating_from: string;
  rating_to: string;
  objetivo_desde: number;
  objetivo_a: number;
}

// 📌 Tipamos correctamente la función para evitar `any`
export function useTablaAcciones() {
  const store = useAccionesStore();
  const acciones = ref<Accion[]>([]); // Ahora el array tiene tipo Accion[]
  const loading = ref<boolean>(true); // Se especifica que loading es un booleano

  // 📌 Definimos los filtros con `Record<string, any>` para mayor flexibilidad
  const filters = ref<Record<string, any>>({
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },
    ticker: { value: null, matchMode: FilterMatchMode.CONTAINS },
    empresa: { value: null, matchMode: FilterMatchMode.CONTAINS },
    corretaje: { value: null, matchMode: FilterMatchMode.CONTAINS },
    accion: { value: null, matchMode: FilterMatchMode.CONTAINS },
    rating_from: { value: null, matchMode: FilterMatchMode.EQUALS },
    rating_to: { value: null, matchMode: FilterMatchMode.EQUALS },
    objetivo_desde: { value: null, matchMode: FilterMatchMode.GREATER_THAN },
    objetivo_a: { value: null, matchMode: FilterMatchMode.LESS_THAN },
  });

  // 📌 Se asegura que la API devuelva datos con la estructura esperada
  onMounted(async () => {
    await store.fetchAcciones();
    acciones.value = store.acciones as Accion[]; // Se hace una conversión a la interfaz definida
    loading.value = false;
  });

  return { acciones, filters, loading };
}
