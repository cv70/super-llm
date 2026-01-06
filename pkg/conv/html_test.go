package conv

import (
	"context"
	"testing"
)

const (
	tableHTML = `<html><body>
<table
	style="margin-right:auto; margin-left:auto; -aw-border-insidev:0.5pt single #000000; border-collapse:collapse"
	class="trs_word_table">
	<tbody>
		<tr style="height:14.15pt" class="firstRow">
			<td rowspan="2"
				style="width:155.98pt; border-top:1.5pt solid #000000; border-right:0.75pt solid #000000; border-bottom:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span
						style="font-family:&#39;Times New Roman&#39;; -aw-import:ignore"> </span>
				</p>
			</td>
			<td colspan="2"
				style="width:163.65pt; border-top:1.5pt solid #000000; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; border-bottom:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">绝对额（亿元）</span></p>
			</td>
			<td colspan="2"
				style="width:160.78pt; border-top:1.5pt solid #000000; border-left:0.75pt solid #000000; border-bottom:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-left:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">比上年同期增长（</span><span
						style="font-family:&#39;Times New Roman&#39;">%</span><span
						style="font-family:宋体">）</span></p>
			</td>
		</tr>

		<tr style="height:14.15pt">
			<td
				style="width:81.35pt; border:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p
					style="margin:0pt 2.85pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">三季度</span></p>
			</td>
			<td
				style="width:81.55pt; border:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p
					style="margin:0pt 2.85pt; text-indent:5.25pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">前三季度</span></p>
			</td>
			<td
				style="width:81.45pt; border:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p
					style="margin:0pt 2.85pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">三季度</span></p>
			</td>
			<td
				style="width:78.58pt; border-top:0.75pt solid #000000; border-left:0.75pt solid #000000; border-bottom:0.75pt solid #000000; vertical-align:middle; -aw-border-bottom:0.25pt single; -aw-border-left:0.5pt single; -aw-border-top:0.25pt single">
				<p
					style="margin:0pt 2.85pt; text-indent:5.25pt; text-align:center; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">前三季度</span></p>
			</td>
		</tr>

		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-top:0.75pt solid #000000; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">GDP</span>
				</p>
			</td>
			<td
				style="width:67.52pt; border-top:0.75pt solid #000000; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">354500</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-top:0.75pt solid #000000; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">1015036</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-top:0.75pt solid #000000; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single; -aw-border-top:0.25pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">4.8</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-top:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-top:0.25pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">5.2</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体; font-weight:bold">第一产业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">26889</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">58061</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">4.0</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">3.8</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体; font-weight:bold">第二产业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">124970</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">364020</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">4.2</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">4.9</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体; font-weight:bold">第三产业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">202641</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">592955</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">5.4</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;; font-weight:bold">5.4</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span
						style="font-family:&#39;Times New Roman&#39;; -aw-import:ignore"> </span>
				</p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:宋体">　</span></p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:宋体">　</span></p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:宋体">　</span></p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:宋体">　</span></p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">农林牧渔业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">28401</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">61626</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.1</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.0</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">工业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">103453</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">306004</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">5.8</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">6.1</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-indent:10.5pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span
						style="font-family:&#39;Times New Roman&#39;">#</span><span
						style="font-family:宋体">制造业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">84866</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">254751</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">6.3</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">6.5</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">建筑业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">22473</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">60683</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">-2.3</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">-0.5</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">批发和零售业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">36389</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">104615</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.9</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">5.6</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">交通运输、仓储和邮政业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">16754</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">46266</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.8</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">5.8</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">住宿和餐饮业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">6916</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">18370</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">3.6</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.6</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">金融业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">26789</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">78348</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">5.2</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.9</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">房地产业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">19862</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">62238</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">-0.2</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">0.6</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">信息传输、软件和信息技术服务业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">15638</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">52758</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">11.7</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">11.2</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">租赁和商务服务业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">15854</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">43669</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">8.6</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">9.2</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td
				style="width:155.98pt; border-right:0.75pt solid #000000; border-bottom:1.5pt solid #000000; vertical-align:middle; -aw-border-right:0.5pt single">
				<p
					style="margin:0pt 2.85pt; text-align:left; line-height:12pt; widows:2; orphans:2">
					<span style="font-family:宋体">其他行业</span></p>
			</td>
			<td
				style="width:67.52pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; border-bottom:1.5pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">61972</span>
				</p>
			</td>
			<td
				style="width:67.72pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; border-bottom:1.5pt solid #000000; padding-right:13.82pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">180458</span>
				</p>
			</td>
			<td
				style="width:59.12pt; border-right:0.75pt solid #000000; border-left:0.75pt solid #000000; border-bottom:1.5pt solid #000000; padding-right:22.32pt; vertical-align:middle; -aw-border-left:0.5pt single; -aw-border-right:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">5.5</span>
				</p>
			</td>
			<td
				style="width:64.38pt; border-left:0.75pt solid #000000; border-bottom:1.5pt solid #000000; padding-right:14.2pt; vertical-align:middle; -aw-border-left:0.5pt single">
				<p style="margin:0pt 2.85pt; text-align:right; line-height:12pt">
					<span style="font-family:&#39;Times New Roman&#39;">4.7</span>
				</p>
			</td>
		</tr>
		<tr style="height:14.15pt">
			<td colspan="5"
				style="width:481.9pt; border-top:1.5pt solid #000000; vertical-align:middle">
				<p style="margin:0pt 2.85pt; text-align:left; line-height:12pt">
					<span style="font-family:楷体">注：</span></p>
				<p style="margin:0pt 2.85pt; text-align:left; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">1.</span><span
						style="font-family:楷体">绝对额按现价计算，增长速度按不变价计算；</span></p>
				<p style="margin:0pt 2.85pt; text-align:left; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">2.</span><span
						style="font-family:楷体">三次产业分类依据国家统计局</span><span
						style="font-family:&#39;Times New Roman&#39;">2018</span><span
						style="font-family:楷体">年修订的《三次产业划分规定》；</span></p>
				<p style="margin:0pt 2.85pt; text-align:left; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">3.</span><span
						style="font-family:楷体">行业分类采用《国民经济行业分类（</span><span
						style="font-family:&#39;Times New Roman&#39;">GB/T
						4754</span><span style="font-family:楷体">－</span><span
						style="font-family:&#39;Times New Roman&#39;">2017</span><span
						style="font-family:楷体">）》；</span></p>
				<p style="margin:0pt 2.85pt; text-align:left; line-height:12pt">
					<span
						style="font-family:&#39;Times New Roman&#39;">4.</span><span
						style="font-family:楷体">本表</span><span
						style="font-family:&#39;Times New Roman&#39;">GDP</span><span
						style="font-family:楷体">总量数据中，有的不等于各产业（行业）之和，是由于数值修约误差所致，未作机械调整。</span>
				</p>
			</td>
		</tr>
	</tbody>
</table>
</body></html>`
	simpleHTML = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Test</title>
  <style>
    body { font-family: Arial; background: #f0f0f0; }
    h1 { color: navy; }
  </style>
</head>
<body>
  <h1>Hello from Go!</h1>
  <p>This HTML was converted to an image using chromedp.</p>
</body>
</html>
`
)

func TestHTML2Markdown(t *testing.T) {
	ctx := context.Background()
	markdown, err := HTML2Markdown(ctx, tableHTML)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("table markdown")
	t.Log(markdown)

	markdown, err = HTML2Markdown(ctx, simpleHTML)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("simple markdown")
	t.Log(markdown)
}
