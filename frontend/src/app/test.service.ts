import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { TransferState, makeStateKey } from '@angular/core';
import { Observable, of, tap } from 'rxjs';

const TESTS_KEY = makeStateKey<any[]>('tests');

export interface TestResponseDto {
  name: string;
  api_endpoint: string;
  response: any;
  status_code: number;
  created_at: string;
}

@Injectable({ providedIn: 'root' })
export class TestService {
  private http = inject(HttpClient);
  private transferState = inject(TransferState);

  getTests(): Observable<any[]> {
    const saved = this.transferState.get(TESTS_KEY, null as any);
    if (saved) {
      return of(saved);
    } else {
      return this.http.get<any[]>('http://localhost:8080/api/tests').pipe(
        tap(data => this.transferState.set(TESTS_KEY, data))
      );
    }
  }
}