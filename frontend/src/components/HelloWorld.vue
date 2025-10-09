<script setup>
import { reactive, watch,ref } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { OfficesList } from '../../wailsjs/go/main/App'
import { LoadZakaz } from '../../wailsjs/go/main/App'
import { Printm } from '../../wailsjs/go/main/App'
const deliveryfrom = ref("pvz");
const deliveryto = ref("pvz");
const poluchatel=ref("yurik");
const phone = ref('')
const isDisabled = ref(true)
const disableINN = ref(false)
const disabletolocation = ref(true)
const selectedItem = ref(null)
const fio = ref('')
const inn = ref('')
const items = ref([])
const tariffs = ref([])
const selectedTariff = ref(null)


const data = reactive({
number:14508,
  city: "",         // значение input
  offices:"",
  results: [],    // список городов
  selectedCity: null,

  officess: [], // список полученных офисов в городе
  selectedOffice: null, // объект выбранного офиса
  officesResults: [], 
  tolocation:null,  // список офисов
})

function selectItem1(item) {
  selectedItem.value = item
}

let debounceTimeout = null

// поиск города
watch(() => data.city, (newVal) => {
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
  data.city = item.full_name
 // data.city = item.full_name
  data.selectedCity = item
  data.results = [] // закрыть список городов
  data.offices=""

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
  alert(JSON.stringify(data.selectedOffice))
}

function changeCity(){
  isDisabled.value=false
}

function onChangePoluchatel() {
  if (poluchatel.value === 'yurik') {
    disableINN.value=false
  } else {
    disableINN.value=true
  }
}

function onChangeDeliveryTo(){
    if (deliveryto.value === 'pvz') {
    disabletolocation.value=true
  } else {
    disabletolocation.value=false
  }

}

function setTariff(code){
    const found = tariffs.value.find(t => t.tariff_code === code)
  if (found) {
    selectedTariff.value = found
  }
}

async function onClick(number) {
  try {
    //alert(deliveryto.value=='pvz')
    items.value=[]
    const res = await LoadZakaz(number)
    fio.value =  res.recipient.name
    if (res.recipient.contragent_type == "LEGAL_ENTITY"){
    poluchatel.value="yurik"
    phone.value=res.recipient.phones[0].number
    inn.value=res.recipient.tin
    items.value.push(...res.packages[0].items)
    tariffs.value.push(...res.tariff_list.tariff_codes)
    data.city=`${res.city.city}`

    if (res.shipment_point!=="" && res.delivery_point!==""){
      deliveryfrom.value="pvz"
      deliveryto.value="pvz"
      let off=res.office_list.find(item =>
      item.code==res.delivery_point)
      data.offices=(off.location.address_full)
      setTariff(res.tariff_code)
    }
    //selectItem(res.city.city)
    //alert(JSON.stringify(res.city.city))
    }
  } catch (err) {
    console.error('Ошибка при получении данных:', err)
  }
  
}
</script>
<template>
  <main>
    <!--
    <div id="result" class="result">{{ data.resultText }}</div>
    -->

    <div class="input-wrapper">
<div class="zakazload">
  <input
   
  v-model.number="data.number"
  class="input"
  type="text"
  placeholder="Введите номер заказа..."
  />

  <button
    :class="computedClass"
    :type="type"
    :disabled="disabled"
    @click="onClick(data.number)"
  >
    <slot>Загрузить данные...</slot> <!-- Содержимое кнопки передаётся через слот -->
  </button>

  </div>




      <!-- выбор города -->
      <div class="city-input-group"> 
      <input
        :disabled="isDisabled"
        v-model="data.city"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите название города..."
      />
          <button class="city-btn" @click="changeCity">
    Изменить...
  </button>
      <div v-if="data.results.length & !isDisabled" class="dropdown">
        <ul>
          <li
            v-for="item in data.results"
            :key="item.city_uuid"
            @click="selectItem(item)"
          >
          <!--  < {{ item.full_name }} - {{item.code}} -->
            {{ item.full_name }}
          </li>
        </ul>
      </div>
      </div>



      </div>

<label-group class="radio-poluchatel">
  <label>
    <input type="radio" value="yurik" v-model="poluchatel" @change="onChangePoluchatel"/>
    Юридическое лицо
  </label>
  <label>
    <input type="radio" value="fizik" v-model="poluchatel" @change="onChangePoluchatel"/>
    Физическое лицо
  </label>
</label-group>

      <div class="tempinput">
      <input
        :disabled="disableINN"
        v-model="inn"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите ИНН организации..."
      />
      <input
        v-model="fio"
        autocomplete="off"
        class="input"
        type="text"
        placeholder="Введите ФИО..."
      />
      <input
        v-model="phone"
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
        v-for="(item,index) in items"
        :key="index"
        class="item"
        :class="{ active: selectedItem === item }"
        @click="selectItem1(item)"
      >
        {{ item.name }}
      </div>
    </div>

    <!-- Редактируемые поля -->
    <div class="inputs-box" v-if="selectedItem">
  <label class="input-group">
    Вес, гр
    <input v-model="selectedItem.weight" type="text" placeholder="Введите вес" />
  </label>

  <label class="input-group">
    Объявленная стоимость
    <input v-model="selectedItem.cost" type="number" placeholder="Введите стоимость" />
  </label>        <label class="input-group">
    Количество
    <input v-model="selectedItem.amount" type="text" placeholder="Введите количество" />
  </label>
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
    <input type="radio" value="pvz" v-model="deliveryto" @change="onChangeDeliveryTo" />
    До ПВЗ
  </label>
  <label>
    <input type="radio" value="cour" v-model="deliveryto" @change="onChangeDeliveryTo" />
    До двери
  </label>
</label-group>
      <!-- выбор офиса -->
      <div class="field"> 
        
                    <input
        v-if="deliveryfrom ==='cour'" 
        :disabled="isDisabled"
        v-model="data.offices"
        autocomplete="off"
        class="input"
        type="text"
         placeholder="Введите адрес забора груза отправителя"
      />


      <input
        :disabled="!(deliveryto === 'pvz' && isDisabled === false)"
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
        :disabled="!(deliveryto === 'cour' && isDisabled === false)"
        v-model="data.tolocation"
        autocomplete="off"
        class="input"
        type="text"
         placeholder="Введите адрес доставки груза получателя"
      />

          <div class="items-box">
      <div
        v-for="(item,index) in tariffs"
        :key="index"
        class="item"
        :class="{ active: selectedTariff === item }"
        @click="selectedTariff = item"
      >
        {{ item.tariff_name }} - {{ item.delivery_sum }}
      </div>
    </div>



    </div>
  </main>
</template>

<style scoped>

.button {
  font-size: 15px;
  padding: 8px 20px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}


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
  color: black;
}

.item {
  padding: 5px;
  margin-bottom: 5px;
  background: rgb(255, 255, 255);
  border: 1px solid #eee;
  border-radius: 3px;
  cursor: pointer;
}

.item:hover {
  background: #beb9b9;
}
.item.active {
  background: #8edb62;
  border-color: #e70cb0;
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

.input-group {
  display: flex;
  flex-direction:column;
  gap: 4px;
  font-weight: 500;
  color: #ffffff;
  
}

.input-group input {
  padding: 4px 6px;
  border: 1px solid #ccc;
  border-radius: 4px;
  text-align: left; 
}
.city-input-group {
  display: flex;
  align-items: center;
  gap: 40px; /* расстояние между полем и кнопкой */
  width:100%;
}

.city-input-group .input {
  flex: 1; /* растягиваем поле, чтобы занимало всё свободное место */
}
</style>