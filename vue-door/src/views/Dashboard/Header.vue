<script setup>
import SearchIcon from "../../assets/icons/IconSearch.vue";
import LogoIcon from "../../assets/icons/IconLogo.vue";
import MenuIcon from "../../assets/icons/IconMenu.vue";
import BasketIcon from "../../assets/icons/IconBasket.vue";
import DropDownIcon from "../../assets/icons/DropdownIcon.vue"
</script>

<script>
export default {
  components:{
    SearchIcon,
    LogoIcon,
    MenuIcon,
    BasketIcon,
    DropDownIcon,
  },

  data() {
    return {
      isDeliverState: true,
    };
  },
  methods: {
    switchDelivery(state) {
      this.isDeliverState = state;
    },
  },
};

</script>


<template>
  <div class="flex flex-row justify-between fixed top-0 left-0 right-0 h-12 w-full border-b-2 border-gray-200">
    <div class="flex flex-row grap-x-1">
      <div class="flex items-center mx-3">
        <MenuIcon />
      </div>
      <div class="flex xl:hidden items-center block  mx-2 cursor-pointer">
        <LogoIcon />
      </div>
      <div class="hidden lg:flex flex-row gap-x-1 items-center ml-3 cursor-pointer">
        <button class="text-[14px]  font-[700] rounded-full px-3 h-8"
          :class="{'bg-black text-white hover:bg-neutral-500': isDeliverState, }"
          @click="switchDelivery(true)"
        >
          Delivery
        </button>  
        <button class="text-[14px]  font-bold rounded-full px-3 h-8"
          :class="{'bg-black text-white hover:bg-neutral-500': !isDeliverState, }"
          @click="switchDelivery(false)"
        >
          Pickup
        </button>  
      </div>
      <div class="hidden lg:block border-l-2 my-3"></div>
      <div class="hidden lg:flex flex-row grap-x-3 items-center">
        <div class="flex flex-row items-center gap-x-0">
          <button class="text-[14px] font-bold  px-3 h-8 flex flex-row items-center"
          >
          ASAP
          <DropDownIcon />
          </button>
          <p> to </p>
          <button class="text-[14px] font-bold  px-3 h-8 flex flex-row items-center"
          >
          Fremont
          <DropDownIcon />
          </button>
        </div>
      </div>
      
    </div>
    
    <div class=" flex-row items-center gap-x-2 hidden xl:flex">
      <div>
        <LogoIcon />
      </div>
      <h2 class="font-bold text-[18px] text-[#ff3008]"> DOORDASH </h2>
    </div>

    <div class="flex flex-row items-center gap-x-4"> 
      <div class="flex flex-row items-center rounded-[20px] bg-[#f7f7f7] border-none px-3 py-[8px] focus-within:shadow-[0_0_0_2px_rgb(0,0,0)] focus-within:bg-white ">
        <SearchIcon />
        <input class=" bg-[#f7f7f7] border-none outline-none focus:bg-white" placeholder="Search" />
      </div>
      <div class="flex flex-row items-center bg-[#eb1700] hover:bg-[#d91400] transition ease-in-out delay-150 duration-300 cursor-pointer text-white px-[15px] mr-[40px] min-h-[32px] rounded-full gap-x-2">
        <BasketIcon />
        <span > 0 </span>
      </div>
    </div>
    
  </div>
</template>