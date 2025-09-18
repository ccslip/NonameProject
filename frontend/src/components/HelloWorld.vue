<script setup>
import { reactive, watch } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { OfficesList } from '../../wailsjs/go/main/App'
import { Prints } from '../../wailsjs/go/main/App'
import { Printm } from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",         // значение input
  offices:"",
  results: [],    // список городов
  selectedCity: null,

  officess: [], // список полученных офисов в городе
  selectedOffice: null, // объект выбранного офиса
  officesResults: [],   // список офисов
  //resultText: "Please enter your city below 👇"
})

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

      <!-- выбор офиса -->
      <input
        v-model="data.offices"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите адрес офиса..."
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
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 30px;
  line-height: 30px;
  margin: 1.5rem auto;
}

.input-wrapper {
  position: relative;
  display: flex;
  flex-direction: column;
  width: 500px;
  margin: 0 auto;
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

.input-wrapper .btn {
  margin-top: 5px;
  height: 30px;
  border-radius: 3px;
  border: none;
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
</style>