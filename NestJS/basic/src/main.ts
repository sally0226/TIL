import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalPipes(
    new ValidationPipe({
      whitelist: true,
      forbidNonWhitelisted: true,
      transform: true, // 받은 데이터를 실제타입으로 변환해줌 ex) string으로 받은 json을 number로
    }),
  );
  await app.listen(3000);
}
bootstrap();
