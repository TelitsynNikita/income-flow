import { inject, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Remain {
  id: number;
  goods_id: number;
  section_id: number;
  goods_count: number;
}

@Injectable( {providedIn: 'root'} )
export class RemainsService {
  private http = inject(HttpClient);
  private url = 'http://localhost:8080'


  getRemains(): Observable<Remain[]> {
    return this.http.get<Remain[]>(`${this.url}/operations/get_remains`);
  }
}
