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
  <div class="lg:w-[70%] w-[90%] relative">
    <button v-if="this.currentScrollState !== ScrollbarState.LEFT"  @click="navigate(true)" class=" flex flex-row items-center justify-center h-[32px] w-[32px] rounded-full bg-white drop-shadow-[0_2px_8px_rgba(0,0,0,0.2)] absolute left-[0px] top-[50%] mt-[-16px] z-10">
      <b-icon-chevron-left class="text-[20px] fill-black cursor-pointer" />
    </button>
    <div ref="container" class="w-full flex flex-row items-center overflow-x-auto scrollbar-hide my-3">
      <slot />
    </div>
    <button v-if="this.currentScrollState !== ScrollbarState.RIGHT"  @click="navigate(false)" class="flex flex-row items-center justify-center h-[32px] w-[32px] rounded-full bg-white drop-shadow-[0_2px_8px_rgba(0,0,0,0.2)] absolute right-[0px] top-[50%] mt-[-16px] z-10">
      <b-icon-chevron-right class="text-[20px] fill-black cursor-pointer" />
    </button>
   
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