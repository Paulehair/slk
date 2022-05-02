<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="movies"
      :items-per-page="100000"
      class="elevation-1"
    ></v-data-table>
  </v-container>
</template>

<script>
export default {
  name: "HelloWorld",

  data: () => ({
    headers: [
      { text: "Titre", value: "title" },
      { text: "Catégorie", value: "category" },
      { text: "Nombre de locations", value: "total_rental" }
    ],
    movies: []
  }),
  mounted() {
    this.$http
      .get("/movies", {
        query: {
          limit: 10000
        }
      })
      .then(res => (this.movies = res.data));
  }
};
</script>
