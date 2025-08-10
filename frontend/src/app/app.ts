import { Component, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon'; // <-- Импорт для иконок
import { MatChipsModule } from '@angular/material/chips'; // <-- Импорт для чипов
import { MatToolbarModule } from '@angular/material/toolbar'; // <-- Импорт для тулбара
import { animate, state, style, transition, trigger } from '@angular/animations'; // <-- Импорт для анимаций

import { TestService, TestResponseDto } from './test.service';

@Component({
  selector: 'app-root',
  standalone: true,
  // Добавляем новые модули в imports
  imports: [
    CommonModule,
    MatTableModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule,
    MatChipsModule,
    MatToolbarModule
  ],
  templateUrl: './app.html',
  styleUrls: ['./app.css'],
  // Добавляем анимацию для раскрытия строки
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0', visibility: 'hidden' })),
      state('expanded', style({ height: '*', visibility: 'visible' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class App implements OnInit {
  // Убираем 'response' из основных колонок, так как это теперь кнопка
  displayedColumns = ['name', 'api_endpoint', 'status_code', 'created_at', 'actions'];
  tests = signal<TestResponseDto[]>([]);
  expandedTest = signal<TestResponseDto | null>(null);

  constructor(private testService: TestService) {}

  ngOnInit() {
    this.testService.getTests().subscribe(data => {
      this.tests.set(data);
    });
  }

  // Логика переключения остается прежней
  toggleExpand(test: TestResponseDto) {
    this.expandedTest.set(this.expandedTest() === test ? null : test);
  }

  // Функция для форматирования JSON
  formatJson(jsonData: any): string {
    try {
      return JSON.stringify(jsonData, null, 2);
    } catch {
      return String(jsonData);
    }
  }

  // Вспомогательная функция для цвета чипа в зависимости от статуса
  getStatusChipColor(statusCode: number): 'primary' | 'accent' | 'warn' {
    if (statusCode >= 200 && statusCode < 300) {
      return 'primary'; // Зеленый (или основной цвет темы) для успешных
    }
    if (statusCode >= 400 && statusCode < 500) {
      return 'accent'; // Оранжевый/желтый для ошибок клиента
    }
    if (statusCode >= 500) {
      return 'warn'; // Красный для ошибок сервера
    }
    return 'primary';
  }
}