import {inject, Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable( {providedIn: 'root'} )
export class OutflowService {
  private http = inject(HttpClient);
  private url = 'http://localhost:8080'


  outflow(goods_id: number, goods_count: number, contractors_id: number, section_id: number) {
    return this.http.post(`${this.url}/operations/outflow`, {
      goods_id,
      goods_count,
      contractors_id,
      section_id,
    });
  }
}
