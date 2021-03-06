#include <cstdint>

#define UNROLL_FACTOR 8

typedef uint8_t byte3 __attribute__((ext_vector_type(3)));
#define ROUND_DOWN(x, s) ((x) & ~((s)-1))

#if 0
extern "C" void hwc2cwh(uint8_t *output, uint8_t *input0, uint64_t height, uint64_t width) {
  const int channels = 3;
  const byte3 *__restrict__ input = (byte3 *)input0;
  uint8_t *const __restrict__ firstOutputPlane = &output[0 * width * height];
  uint8_t *const __restrict__ secondOutputPlane = &output[1 * width * height];
  uint8_t *const __restrict__ thirdOutputPlane = &output[2 * width * height];
  int offset = 0;
  for (uint64_t yy = 0; yy < height; yy++) {
    uint64_t xx = 0;
    for (xx = 0; xx < width; xx++) {
      const float3 rgb = input[offset];
      firstOutputPlane[offset] = rgb.x;
      secondOutputPlane[offset] = rgb.y;
      thirdOutputPlane[offset] = rgb.z;
      offset++;
    }
  }
}
#else
extern "C" void hwc2chw(uint8_t *__restrict__ output,
                        const uint8_t *const __restrict__ input0,
                        const uint64_t height, const uint64_t width) {
  const int channels = 3;
  const byte3 *__restrict__ input = (byte3 *)input0;
  uint8_t *const __restrict__ firstOutputPlane = &output[0 * width * height];
  uint8_t *const __restrict__ secondOutputPlane = &output[1 * width * height];
  uint8_t *const __restrict__ thirdOutputPlane = &output[2 * width * height];
  for (uint64_t yy = 0; yy < height; yy++) {
    uint64_t xx = 0;
    for (; xx < ROUND_DOWN(width, UNROLL_FACTOR); xx += UNROLL_FACTOR) {
      const uint64_t offset = yy * width + xx;
      __builtin_prefetch(&input[(offset + 1) * UNROLL_FACTOR], 0, 1);
#pragma unroll
      for (int ii = 0; ii < UNROLL_FACTOR; ii++) {
        const byte3 rgb = input[offset + ii];
        firstOutputPlane[offset + ii] = rgb.x;
        secondOutputPlane[offset + ii] = rgb.y;
        thirdOutputPlane[offset + ii] = rgb.z;
      }
    }

    for (; xx < width; xx++) {
      const uint64_t offset = yy * width + xx;
      const byte3 rgb = input[offset];
      firstOutputPlane[offset] = rgb.x;
      secondOutputPlane[offset] = rgb.y;
      thirdOutputPlane[offset] = rgb.z;
    }
  }
}
#endif
