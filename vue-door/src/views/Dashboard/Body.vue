<script>
  import NavigatorItem from '../../components/NavigatorItem.vue'
  import navigatorData from '../../assets/data/navigate.json'
  import Categoryitem from '../../components/CategoryItem.vue'
  import SliderDiv from '../../components/SliderDiv.vue'
  import DropdownButton from '../../components/DropdownButton.vue'
  import ToggleButton from '../../components/ToggleButton.vue'
  import PassIcon from '../../assets/icons/PassIcon.vue'
  import NormalButton from '../../components/NormalButton.vue'
  import SliderContaier from '../../components/SliderContainer.vue'
  export default {
    components:{
      NavigatorItem,
      SliderDiv,
      Categoryitem,
      DropdownButton,
      ToggleButton,
      PassIcon,
      NormalButton,
      SliderContaier,
    },
    data() {
      return {
        navigatorData: navigatorData,
        sliderValue: 5,
      }
    },
    methods: {
      handleRangeChange($event){
        console.log(event.target.value)
        this.sliderValue = event.target.value
      }
    }

  }
</script>

<template>
  <div class="top-12 flex flex-col w-full items-center">
    <div class="flex flex-row items-center bg-[#f7f7f7]  w-full justify-center ">
      <slider-div >
        <NavigatorItem v-for="navigator in navigatorData" :key="navigator.title" :title="navigator.title" :imgUrl="navigator.url"  ></NavigatorItem>
      </slider-div>
    </div> 
    <div class="flex flex-row items-center w-full justify-center mt-[-9px] mb-[10px]">
      <slider-div >
        <Categoryitem v-for="item in [1,2,3,5,5,5,5,5,55,5,5,5,5,5,55]" :key="item" />
      </slider-div>
    </div>
    <div class="flex flex-row items-center lg:w-[70%] w-[90%] relative justify-center ">
      <div class="flex flex-row items-center w-full justify-start gap-2 ">
        <ToggleButton>Pickup</ToggleButton>
        <dropdown-button>
          <template v-slot:button>
            <div class="content flex flex-row items-center">
              <p class="font-bold"> Over 4.5 </p>
              <b-icon-star-fill class="fill-black pl-1" />
            </div>
          </template>
          <template v-slot:content>
            <div class="content flex flex-col ">
              <p class="content text-[24px] mt-[10px] mx-4">Rating</p>
              <p class="content text-[16px] mx-4">Over {{sliderValue}}</p>
              
              <div class="mx-4 mt-4 ">
                <div class="range relative w-[300px] h-[5px] flex flex-row items-center justify-center bg-gradient-to-l from-[#000] to-[#b1b1b1]" >
                  <input type="range" v-on:input="handleRangeChange($event)" :data="sliderValue"  min="3" max="5" 
                    class="bg-gradient-to-l from-[#000]  to-[#b1b1b1] w-full absolute t-[2px] h-[2px] bg-[#b2b2b2] appearance-none z-[10] " step="0.5" 
                    
                  />
                 </div>
                <ul class="range-labels ">
                  <li :class="{'selected': sliderValue <= 3}">3</li>
                  <li :class="{'selected': sliderValue <= 3.5}">3.5</li>
                  <li :class="{'selected': sliderValue <= 4}">4</li>
                  <li :class="{'selected': sliderValue <= 4.5}">4.5</li>
                  <li :class="{'selected': sliderValue <= 5}">5</li>
                 </ul>
              </div>
              
              <hr class="mt-2"/>
              <div class="flex flex-row justify-end gap-5">
                <button>cancel </button>
                <normal-button class="my-4 mx-4"> View Results </normal-button>
              </div>
            </div>
          </template>
        </dropdown-button>
        <ToggleButton>Under 30min</ToggleButton>
        <dropdown-button>
          <template v-slot:button>
            <div class="content flex flex-row items-center">
              <p class="font-bold"> Price </p>
            </div>
          </template>
          <template v-slot:content>
            <div class="content flex flex-col p-2">
              <p class="content text-[24px]">Price</p>
            </div>
          </template>
        </dropdown-button>
        <ToggleButton>
          <pass-icon />
          DashPass
        </ToggleButton>
      </div>
    </div >
    
    <div class="flex flex-row items-center w-full justify-center mt-[20px] mb-[10px]">
      <slider-div >
        <div class="flex flex-row justify-between ml-[2%] min-w-[48%] bg-[#ffe0db] rounded-[8px] min-h-[156px] items-center">
          <div class="max-h-[156px] min-w-[400px] flex flex-col gap-3 px-6">
            <span class="text-[20px] font-bold">
              Get $0 delivery fee on your<br /> first over $10
            </span>
            <button class="font-bold py-1 w-[120px] rounded-full bg-[#ffc4be]"> Learn More </button>
          </div>
          <div class="grow p-0">
                <img class="h-[156px]  rounded-r-[8px]  w-[200px]" src="@/assets/img/2.png" />
          </div>
        </div>
        <div class="flex flex-row justify-between ml-[2%] min-w-[48%] bg-[#ee6f68] rounded-[8px] min-h-[156px] items-center">
          <div class="max-h-[156px] min-w-[400px] flex flex-col gap-2 px-5 text-white">
            <span class="text-[20px]  font-bold">
              Shop great deals from your <br /> favourite local spots
            </span>
            <span>check out our Offers Hub!</span>
            <button class="font-bold py-1 w-[120px] rounded-full bg-[#f4a09c]"> Order Now </button>              
          </div>
          <div class=" p-0">
                <img class="h-[156px]  rounded-r-[8px]  w-[200px]" src="@/assets/img/2.png" />
          </div>
        </div>
        <div class="flex flex-row justify-between ml-[2%] min-w-[48%] bg-[#ffe0db] rounded-[8px] min-h-[156px] items-center">
          <div class="max-h-[156px] min-w-[400px] flex flex-col gap-3 px-6">
            <span class="text-[20px] font-bold">
              Get $0 delivery fee on your<br /> first over $10
            </span>
            <button class="font-bold py-1 w-[120px] rounded-full bg-[#ffc4be]"> Learn More </button>
          </div>
          <div class="p-0">
                <img class="h-[156px]  rounded-r-[8px]  w-[200px]" src="@/assets/img/2.png" />
          </div>
        </div>
      </slider-div>
    </div>
    
    <div class="flex items-center w-full justify-center">
      <slider-contaier>
        <template #title>
          <p class="text-[24px] font-bold">Fastest Near You</p>
        </template>
        <template #content>
          <div class="min-w-[30%] flex flex-col">
            <img class="h-[172px] rounded-[6px]" 
              src="https://img.cdn4dd.com/p/fit=cover,width=400,format=jpeg,quality=50/media/photos/40a20c10-de38-4c47-adbe-1353529f5f94-retina-large-jpeg" />
            <div class="flex flex-row justify-between mt-3">
              <div class="flex flex-row gap-1">
                <p class="text-[16px] font-bold"> Taco Bell</p>
                <pass-icon />
              </div>
              <b-icon-heart class=" fill-black text-[18px] mr-2"/>
            </div>
            <p> 18min*2.1mi*$0 delivery fee</p>
            <div class="flex items-center flex-row gap-1">
              <p>4.3 </p>
              <b-icon-star-fill class="fill-black mr-1 text-[10px]" />
              <p>(1900+)</p>
            </div>
          </div>
          <div class="min-w-[30%] flex flex-col">
            <img class="h-[172px] rounded-[6px]" 
              src="https://img.cdn4dd.com/p/fit=cover,width=400,format=jpeg,quality=50/media/photosV2/8495f9d8-c7c5-46de-985f-0d9c214dd4f6-retina-large.jpg" />
            <div class="flex flex-row justify-between mt-3">
              <div class="flex flex-row gap-1">
                <p class="text-[16px] font-bold"> Taco Bell</p>
                <pass-icon />
              </div>
              <b-icon-heart class=" fill-black text-[18px] mr-2"/>
            </div>
            <p> 18min*2.1mi*$0 delivery fee</p>
            <div class="flex items-center flex-row gap-1">
              <p>4.3 </p>
              <b-icon-star-fill class="fill-black mr-1 text-[10px]" />
              <p>(1900+)</p>
            </div>
          </div>
          <div class="min-w-[30%] flex flex-col">
            <img class="h-[172px] rounded-[6px]" 
              src="https://img.cdn4dd.com/p/fit=cover,width=400,format=jpeg,quality=50/media/store/header/169b1ffa-fc81-418b-8038-120096972272.jpeg" />
            <div class="flex flex-row justify-between mt-3">
              <div class="flex flex-row gap-1">
                <p class="text-[16px] font-bold"> Taco Bell</p>
                <pass-icon />
              </div>
              <b-icon-heart class=" fill-black text-[18px] mr-2"/>
            </div>
            <p> 18min*2.1mi*$0 delivery fee</p>
            <div class="flex items-center flex-row gap-1">
              <p>4.3 </p>
              <b-icon-star-fill class="fill-black mr-1 text-[10px]" />
              <p>(1900+)</p>
            </div>
          </div>
          <div class="min-w-[30%] flex flex-col">
            <img class="h-[172px] rounded-[6px]" 
              src="https://img.cdn4dd.com/p/fit=cover,width=400,format=jpeg,quality=50/media/photos/98a6f342-aa36-453e-ab97-33e377f52cfb-retina-large-png" />
            <div class="flex flex-row justify-between mt-3">
              <div class="flex flex-row gap-1">
                <p class="text-[16px] font-bold"> Taco Bell</p>
                <pass-icon />
              </div>
              <b-icon-heart class=" fill-black text-[18px] mr-2"/>
            </div>
            <p> 18min*2.1mi*$0 delivery fee</p>
            <div class="flex items-center flex-row gap-1">
              <p>4.3 </p>
              <b-icon-star-fill class="fill-black mr-1 text-[10px]" />
              <p>(1900+)</p>
            </div>
          </div>
        </template>
      </slider-contaier>
    </div>
      

  </div>
</template>
<style scoped>
input[type=range]::-webkit-slider-thumb {
  -webkit-appearance: none;
  height: 20px;
  width: 20px;
  border-radius: 10px;
  background: black;
  cursor: pointer;
}

.range-labels {
  margin-top: 20px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: 0;
  list-style: none;
}

.range-labels li {
    position: relative;
    float: left;
    text-align: center;
    color: #b2b2b2;
    font-size: 14px;
    cursor: pointer;
}
.range-labels li::before {
      position: absolute;
      top: -27px;
      right: 0;
      left: 0;
      content: "";
      margin: 0 auto;
      width: 9px;
      height: 9px;
      background: #b2b2b2;
      border-radius: 50%;
    }
  
  .active {
    color: #37adbf;
  }
  
  li.selected::before {
    background: black;
  }
  
  .active.selected::before {
    display: none;
  }
  /* .range {background: linear-gradient(to left, black 0%, black 75%, #fff 75%, #fff 100%)}
  .range input{background: linear-gradient(to left, black 0%, black 75%, #b2b2b2 75%, #b2b2b2 100%)} */
</style>
