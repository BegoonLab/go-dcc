<template>
      <p>
          {{ message }}
      </p>
</template>

<script>
export default {
  data() {
    return {
      message: 'Hello World',
    };
  },
};
</script>

<style>
p {
  font-size: 18px;
  font-family: 'Roboto', sans-serif;
  color: blue;
}
</style>