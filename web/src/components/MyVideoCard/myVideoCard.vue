// components/VideoCard.vue
<template>
  <el-card class="my-video-card" shadow="hover">
    {{props.title}}
    <el-row class="my-video-card-content" :gutter="20">
      <!-- 左侧图片 -->
      <el-col :span="6" class="image-container">
        <img :src="imageSrc" alt="image" class="video-image" />
      </el-col>

      <!-- 右侧内容 -->
      <el-col :span="18" class="content-container">
        <div class="header-section">
          <!-- 标题 -->
          <h2 class="video-title">{{ title }}</h2>
          <!-- 评分 -->
          <div class="video-score">评分：{{ score }}</div>
        </div>
        <!-- 介绍 -->
        <div class="video-description">
          <strong>介绍：</strong>
          <span>{{ description }}</span>
        </div>
        <!-- 热评 -->
        <div class="video-hotcomment">
          <strong>热评：</strong>
          <span>{{ hotComment }}</span>
        </div>
        <!-- 按钮链接 -->
        <div class="video-links">
          <el-button v-for="(url, index) in urls" :key="index" round size="small" @click="handleLink(url)">
            url
          </el-button>
        </div>
      </el-col>
    </el-row>
  </el-card>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

// 定义 Props
const props = defineProps({
  imageSrc: {
    type: String,
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
  score: {
    type: [Number, String],
    default: 'N/A',
  },
  description: {
    type: String,
    default: '',
  },
  hotComment: {
    type: String,
    default: '',
  },
  urls: {
    type: Array as () => string[],
    default: () => [],
  },
});

// 方法：打开链接
const handleLink = (url: string) => {
  window.open(url, '_blank');
};

console.log(props)
</script>

<style scoped lang="scss">
@import "./static/style/myVideoCard";
</style>