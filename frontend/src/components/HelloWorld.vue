<script setup>
import { reactive, watch,ref } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { OfficesList } from '../../wailsjs/go/main/App'
import { Prints } from '../../wailsjs/go/main/App'
import { Printm } from '../../wailsjs/go/main/App'
const deliveryfrom = ref("pvz");
const deliveryto = ref("pvz");
const poluchatel=ref("yurik");
const data = reactive({
  inn:"",
  phone:"",
  fio:"",
  name: "",         // значение input
  offices:"",
  results: [],    // список городов
  selectedCity: null,

  officess: [], // список полученных офисов в городе
  selectedOffice: null, // объект выбранного офиса
  officesResults: [],   // список офисов
})
const items = ref([
  { name: "Товар 1", price: 100, sku: "A001" },
  { name: "Товар 2", price: 200, sku: "A002" },
  { name: "Товар 3", price: 300, sku: "A003" }
])

const selectedItem = ref(null)

function selectItem1(item) {
  selectedItem.value = item
}

let debounceTimeout = null

// поиск города
watch(() => data.name, (newVal) => {
  if (debounceTimeout) clearTimeout(debounceTimeout)

  debounceTimeout = setTimeout(async () => {
    if (newVal.trim() === "") {
      data.results = []
      return
    }

    try {
      const res = await Greet(newVal)
      // проверяем, что вернулся массив объектов
      data.results = Array.isArray(res) ? res : []
    } catch (e) {
      console.error(e)
      data.results = []
    }
  }, 300) // debounce 300ms
})

//поиск адреса в списке data.offices
watch(() => data.offices, (newVal) => {
  if (!data.selectedCity) {
  return
  }
  if (data.selectedOffice && newVal === data.selectedOffice.location.address_full) {
  return
  }


  if (debounceTimeout) clearTimeout(debounceTimeout)

  debounceTimeout = setTimeout(async () => {
    if (newVal.trim() === "") {
      data.officesResults = []
      return
    }
    //Prints(data.officesResults)
    const lowerQuery = newVal.toLowerCase()
    data.officesResults = data.officess.filter(item =>
      item.location.address_full.toLowerCase().includes(lowerQuery)
    )
  }, 300) // debounce 100ms
})


async function selectItem(item) {
  data.name = item.full_name
  data.selectedCity = item
  data.results = [] // закрыть список городов
  //data.offices = OfficesList(item.code.toString())

  try {
    // загружаем список офисов для выбранного города
    const offices = await OfficesList(item.code.toString())
    data.officess = Array.isArray(offices) ? offices : []
    data.officesResults = data.officess   // сразу показываем все офисы
  } catch (e) {
    console.error(e)
    data.officess = []
    data.officesResults = []
  }
//  try {
    // подгрузка офисов по коду города
//    const offices = await OfficesList(item.code.toString())
//    data.officesResults = Array.isArray(offices) ? offices : []
 // } catch (e) {
  //  console.error(e)
  //  data.officesResults = []
 // }
}


function selectOffice(office) {
  data.offices = office.location.address_full   // вставляем название офиса в input
  data.selectedOffice = office // сохраняем объект офиса
  data.officesResults = []     // закрыть список офисов
}

</script>
<template>
  <main>
    <!--
    <div id="result" class="result">{{ data.resultText }}</div>
    -->
    <div class="input-wrapper">
      <!-- выбор города -->
      <div class="field"> 
      <input
        v-model="data.name"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите название города..."
      />

      <div v-if="data.results.length" class="dropdown">
        <ul>
          <li
            v-for="item in data.results"
            :key="item.city_uuid"
            @click="selectItem(item)"
          >
            {{ item.full_name }} - {{item.code}}
          </li>
        </ul>
      </div>
      </div>



      </div>

<label-group class="radio-poluchatel">
  <label>
    <input type="radio" value="yurik" v-model="poluchatel" />
    Юридическое лицо
  </label>
  <label>
    <input type="radio" value="fizik" v-model="poluchatel" />
    Физическое лицо
  </label>
</label-group>

      <div class="tempinput">
      <input
      v-if="poluchatel ==='yurik'"
        v-model="data.inn"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите ИНН организации..."
      />
      <input
        v-model="data.fio"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите ФИО..."
      />
      <input
        v-model="data.phone"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите контактный телефон..."
      />
      </div>

  <div class="wrapper">
    <!-- Список элементов -->
    <div class="items-box">
      <div
        v-for="(item, index) in items"
        :key="index"
        class="item"
        @click="selectItem1(item)"
      >
        {{ item.name }}
      </div>
    </div>

    <!-- Редактируемые поля -->
    <div class="inputs-box" v-if="selectedItem">
      <input v-model="selectedItem.name" type="text" placeholder="Название товара" />
      <input v-model="selectedItem.price" type="number" placeholder="Цена" />
      <input v-model="selectedItem.sku" type="text" placeholder="Артикул" />
    </div>
  </div>

<label-group class="radio-from">
  <label>
    <input type="radio" value="pvz" v-model="deliveryfrom" />
    От ПВЗ
  </label>
  <label>
    <input type="radio" value="cour" v-model="deliveryfrom" />
    От двери
  </label>
</label-group>

<label-group class="radio-from">
  <label>
    <input type="radio" value="pvz" v-model="deliveryto" />
    До ПВЗ
  </label>
  <label>
    <input type="radio" value="cour" v-model="deliveryto" />
    До двери
  </label>
</label-group>
      <!-- выбор офиса -->
      <div class="field"> 
        
                    <input
        v-if="deliveryfrom ==='cour'" 
        v-model="data.offices"
        autocomplete="off"
        class="input"
        type="text"
         placeholder="Введите адрес забора груза отправителя"
      />


      <input
        v-if="deliveryto ==='pvz'" 
        v-model="data.offices"
        autocomplete="off"
        class="input"
        type="text"
         placeholder="Введите адрес ПВЗ получателя"
      />

      <div v-if="data.officesResults.length" class="dropdown">
        <ul>
          <li
            v-for="office in data.officesResults"
            :key="office.uuid"
            @click="selectOffice(office)"
          >
            {{ office.location.address_full }}
          </li>
        </ul>
      </div>

                  <input
        v-if="deliveryto ==='cour'" 
        v-model="data.offices"
        autocomplete="off"
        class="input"
        type="text"
         placeholder="Введите адрес доставки груза получателя"
      />



    </div>
  </main>
</template>

<style scoped>

.field {
  position: relative; 
  display: flex;
  flex-direction: column;
  width: 500px;
  margin: 5px 5px;

}


.input-wrapper {
  position: relative;
  display: flex;
  flex-direction: column;
  width: 500px;
  margin: 5px 5px;
}

.tempinput {
  position: relative;
  display: flex;
  flex-direction: column;
  width: 500px;

}

.input-wrapper .input {
  height: 30px;
  padding: 0 10px;
  border-radius: 3px;
  border: none;
  outline: none;
  font-size: 14px;
  margin-bottom: 5px;
}
.radio-from {
  display: flex;        /* выстраиваем в одну линию */
  gap: 1rem;            /* расстояние между кнопками */
  margin: 5px 0;        /* небольшой отступ сверху/снизу */
  align-items: center;  /* выравниваем по вертикали */
}

.radio-from label {
  cursor: pointer;
}
.radio-poluchatel {
  display: flex;        /* выстраиваем в одну линию */
  gap: 1rem;            /* расстояние между кнопками */
  margin: 20px 0;        /* небольшой отступ сверху/снизу */
  align-items: center;  /* выравниваем по вертикали */
}
.radio-poluchatel label {
  cursor: pointer;
}


.dropdown {
  position: absolute;
  top: 100%;
  left: 0;
  width: 100%;
  max-height: 150px;
  overflow-y: auto;
  border: 1px solid #ccc;
  background: white;
  border-radius: 3px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.15);
  z-index: 10;
}

.dropdown ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.dropdown li {
  padding: 8px 10px;
  cursor: pointer;
   color: black;
}

.dropdown li:hover {
  background: #f0f0f0;
   color: black;
}
.wrapper {
  display: flex;
  gap: 20px; /* расстояние между колонками */
}

.items-box {
  width: 200px;
  height: 200px;
  border: 1px solid #ccc;
  overflow-y: auto;
  padding: 10px;
  background: #fafafa;
}

.item {
  padding: 5px;
  margin-bottom: 5px;
  background: white;
  border: 1px solid #eee;
  border-radius: 3px;
  cursor: pointer;
}

.item:hover {
  background: #eaeaea;
}

.inputs-box {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.inputs-box input {
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 3px;
}



</style>