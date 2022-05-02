<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="movies"
      :items-per-page="15"
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
        params: {
          limit: 200
        }
      })
      .then(res => (this.movies = res.data));

    this.$http.get("/count_pages").then(res => console.log(res.data));
  }
};
</script>
