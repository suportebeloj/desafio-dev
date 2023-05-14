<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="header">
      <div class="header__navbrand">
        desafio dev
      </div>
    </div>
  </header>

  <RouterView />
</template>

<style scoped lang="sass">
$header-color: #f0f0f0
$text-color: #333

.header
  width: 100%
  height: 50px
  background-color: $header-color
  display: flex
  justify-content: space-between
  align-items: center
  padding-inline: 20px

  &__navbrand
    font-size: 1.3em
    font-weight: bold
    font-family: 'Courier New', Courier, monospace
    text-transform: capitalize
    cursor: pointer
    transition: color 300ms

    &:hover
      color: #ccc

.header
  max-width: 100vw
  max-height: 50px
</style>
