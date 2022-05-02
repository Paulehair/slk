<template>
  <v-app>
    <v-app-bar app color="orange" dark>
      <v-title dark>Salika Pagination</v-title>
      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
      <Pagination />
    </v-main>
  </v-app>
</template>

<script>
import Pagination from "./components/Pagination";

export default {
  name: "App",

  components: {
    Pagination
  }
};
</script>
