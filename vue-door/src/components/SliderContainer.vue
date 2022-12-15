<script>
import { ref, onUpdated } from 'vue'

const ScrollbarState = {
  LEFT: 0,
  MIDDLE: 1,
  RIGHT: 2,
};

export default {
    data() {
      return {
      scrollbarPostion: 0,
      currentScrollState: ScrollbarState.LEFT,
      ScrollbarState: ScrollbarState
      }
    },
    methods: {
      navigate(state) {
        // console.log(this.divId)
        const container = this.$refs.container;
        //  console.log()
        if(state){
          container.scrollLeft -= container.clientWidth
          this.scrollbarPostion -= container.clientWidth 
          if(this.scrollbarPostion <= 0){
            this.currentScrollState = ScrollbarState.LEFT
            this.scroll = 0
          }
          else{
             this.currentScrollState = ScrollbarState.MIDDLE
          }
        }
        else if(!state){
          container.scrollLeft += container.clientWidth
          this.scrollbarPostion += container.clientWidth 
          if(this.scrollbarPostion + container.clientWidth >= container.scrollWidth){
            this.scroll = container.scrollWidth - container.clientWidth
            this.currentScrollState = ScrollbarState.RIGHT
          }
          else{
             this.currentScrollState = ScrollbarState.MIDDLE
          }
        }
      },


    },

    async updated() {
      // console.log(this.scroll)
    }
}
</script>

<template>
  <div class="lg:w-[70%] w-[90%] flex flex-col">
    <div class="flex flex-row">
      <div class="flex-1">
        <slot name="title" />
      </div>
      <button class="pr-10 font-bold"> See All </button>
      <button   
        @click="navigate(true)" 
        class="flex flex-row items-center justify-center mr-3 bg-[#e7e7e7] h-[32px] w-[32px] rounded-full  drop-shadow-[0_0px_1px_rgba(0,0,0,0.2)] disabled:opacity-50"
        :disabled="this.currentScrollState == ScrollbarState.LEFT"
        >
        <b-icon-chevron-left class="text-[20px] fill-black cursor-pointer" />
      </button>
      <button 
        @click="navigate(false)" 
        class="flex flex-row items-center justify-center bg-[#e7e7e7] h-[32px] w-[32px] rounded-full  drop-shadow-[0_0px_1px_rgba(0,0,0,0.2)] disabled:opacity-50"
        :disabled="this.currentScrollState == ScrollbarState.RIGHT"
        >
        <b-icon-chevron-right class="text-[20px] fill-black cursor-pointer" />
      </button>
    </div>  
    <div ref="container" class="w-full flex flex-row items-center overflow-x-auto scrollbar-hide gap-[5%] my-3">
      <slot name="content"/>
    </div>
  </div>
</template>

<style scoped>
.scrollbar-hide::-webkit-scrollbar {
  display: none;
  scroll-behavior: smooth;
}

.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;  
  scroll-behavior: smooth;
}
</style>